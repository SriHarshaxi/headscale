package headscale

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"tailscale.com/tailcfg"
	"tailscale.com/types/wgkey"
)

// PollNetMapHandler takes care of /machine/:id/map
//
// This is the busiest endpoint, as it keeps the HTTP long poll that updates
// the clients when something in the network changes.
//
// The clients POST stuff like HostInfo and their Endpoints here, but
// only after their first request (marked with the ReadOnly field).
//
// At this moment the updates are sent in a quite horrendous way, but they kinda work.
func (h *Headscale) PollNetMapHandler(c *gin.Context) {
	log.Trace().
		Str("handler", "PollNetMap").
		Str("id", c.Param("id")).
		Msg("PollNetMapHandler called")
	body, _ := io.ReadAll(c.Request.Body)
	mKeyStr := c.Param("id")
	mKey, err := wgkey.ParseHex(mKeyStr)
	if err != nil {
		log.Error().
			Str("handler", "PollNetMap").
			Err(err).
			Msg("Cannot parse client key")
		c.String(http.StatusBadRequest, "")
		return
	}
	req := tailcfg.MapRequest{}
	err = decode(body, &req, &mKey, h.privateKey)
	if err != nil {
		log.Error().
			Str("handler", "PollNetMap").
			Err(err).
			Msg("Cannot decode message")
		c.String(http.StatusBadRequest, "")
		return
	}

	var m Machine
	if result := h.db.Preload("Namespace").First(&m, "machine_key = ?", mKey.HexString()); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Warn().
			Str("handler", "PollNetMap").
			Msgf("Ignoring request, cannot find machine with key %s", mKey.HexString())
		c.String(http.StatusUnauthorized, "")
		return
	}
	log.Trace().
		Str("handler", "PollNetMap").
		Str("id", c.Param("id")).
		Str("machine", m.Name).
		Msg("Found machine in database")

	hostinfo, _ := json.Marshal(req.Hostinfo)
	m.Name = req.Hostinfo.Hostname
	m.HostInfo = datatypes.JSON(hostinfo)
	m.DiscoKey = wgkey.Key(req.DiscoKey).HexString()
	now := time.Now().UTC()

	// From Tailscale client:
	//
	// ReadOnly is whether the client just wants to fetch the MapResponse,
	// without updating their Endpoints. The Endpoints field will be ignored and
	// LastSeen will not be updated and peers will not be notified of changes.
	//
	// The intended use is for clients to discover the DERP map at start-up
	// before their first real endpoint update.
	if !req.ReadOnly {
		endpoints, _ := json.Marshal(req.Endpoints)
		m.Endpoints = datatypes.JSON(endpoints)
		m.LastSeen = &now
	}
	h.db.Save(&m)

	data, err := h.getMapResponse(mKey, req, m)
	if err != nil {
		log.Error().
			Str("handler", "PollNetMap").
			Str("id", c.Param("id")).
			Str("machine", m.Name).
			Err(err).
			Msg("Failed to get Map response")
		c.String(http.StatusInternalServerError, ":(")
		return
	}

	// We update our peers if the client is not sending ReadOnly in the MapRequest
	// so we don't distribute its initial request (it comes with
	// empty endpoints to peers)

	// Details on the protocol can be found in https://github.com/tailscale/tailscale/blob/main/tailcfg/tailcfg.go#L696
	log.Debug().
		Str("handler", "PollNetMap").
		Str("id", c.Param("id")).
		Str("machine", m.Name).
		Bool("readOnly", req.ReadOnly).
		Bool("omitPeers", req.OmitPeers).
		Bool("stream", req.Stream).
		Msg("Client map request processed")

	if req.ReadOnly {
		log.Info().
			Str("handler", "PollNetMap").
			Str("machine", m.Name).
			Msg("Client is starting up. Probably interested in a DERP map")
		c.Data(200, "application/json; charset=utf-8", *data)
		return
	}

	// There has been an update to _any_ of the nodes that the other nodes would
	// need to know about
	h.setLastStateChangeToNow()

	// The request is not ReadOnly, so we need to set up channels for updating
	// peers via longpoll

	// Only create update channel if it has not been created
	log.Trace().
		Str("handler", "PollNetMap").
		Str("id", c.Param("id")).
		Str("machine", m.Name).
		Msg("Loading or creating update channel")
	var updateChan chan struct{}
	if storedChan, ok := h.clientsUpdateChannels.Load(m.ID); ok {
		if wrapped, ok := storedChan.(chan struct{}); ok {
			updateChan = wrapped
		} else {
			log.Error().
				Str("handler", "PollNetMap").
				Str("id", c.Param("id")).
				Str("machine", m.Name).
				Msg("Failed to convert update channel to struct{}")
		}
	} else {
		log.Debug().
			Str("handler", "PollNetMap").
			Str("id", c.Param("id")).
			Str("machine", m.Name).
			Msg("Update channel not found, creating")

		updateChan = make(chan struct{})
		h.clientsUpdateChannels.Store(m.ID, updateChan)
	}

	pollDataChan := make(chan []byte)
	// defer close(pollData)

	keepAliveChan := make(chan []byte)

	cancelKeepAlive := make(chan struct{})
	defer close(cancelKeepAlive)

	if req.OmitPeers && !req.Stream {
		log.Info().
			Str("handler", "PollNetMap").
			Str("machine", m.Name).
			Msg("Client sent endpoint update and is ok with a response without peer list")
		c.Data(200, "application/json; charset=utf-8", *data)

		// It sounds like we should update the nodes when we have received a endpoint update
		// even tho the comments in the tailscale code dont explicitly say so.
		go h.notifyChangesToPeers(&m)
		return
	} else if req.OmitPeers && req.Stream {
		log.Warn().
			Str("handler", "PollNetMap").
			Str("machine", m.Name).
			Msg("Ignoring request, don't know how to handle it")
		c.String(http.StatusBadRequest, "")
		return
	}

	log.Info().
		Str("handler", "PollNetMap").
		Str("machine", m.Name).
		Msg("Client is ready to access the tailnet")
	log.Info().
		Str("handler", "PollNetMap").
		Str("machine", m.Name).
		Msg("Sending initial map")
	go func() { pollDataChan <- *data }()

	log.Info().
		Str("handler", "PollNetMap").
		Str("machine", m.Name).
		Msg("Notifying peers")
	go h.notifyChangesToPeers(&m)

	h.PollNetMapStream(c, m, req, mKey, pollDataChan, keepAliveChan, updateChan, cancelKeepAlive)
	log.Trace().
		Str("handler", "PollNetMap").
		Str("id", c.Param("id")).
		Str("machine", m.Name).
		Msg("Finished stream, closing PollNetMap session")
}

func (h *Headscale) PollNetMapStream(
	c *gin.Context,
	m Machine,
	req tailcfg.MapRequest,
	mKey wgkey.Key,
	pollDataChan chan []byte,
	keepAliveChan chan []byte,
	updateChan chan struct{},
	cancelKeepAlive chan struct{},
) {
	go h.keepAlive(cancelKeepAlive, keepAliveChan, mKey, req, m)

	c.Stream(func(w io.Writer) bool {
		log.Trace().
			Str("handler", "PollNetMapStream").
			Str("machine", m.Name).
			Msg("Waiting for data to stream...")

		log.Trace().
			Str("handler", "PollNetMapStream").
			Str("machine", m.Name).
			Msgf("pollData is %#v, keepAliveChan is %#v, updateChan is %#v", pollDataChan, keepAliveChan, updateChan)

		select {
		case data := <-pollDataChan:
			log.Trace().
				Str("handler", "PollNetMapStream").
				Str("machine", m.Name).
				Str("channel", "pollData").
				Int("bytes", len(data)).
				Msg("Sending data received via pollData channel")
			_, err := w.Write(data)
			if err != nil {
				log.Error().
					Str("handler", "PollNetMapStream").
					Str("machine", m.Name).
					Str("channel", "pollData").
					Err(err).
					Msg("Cannot write data")
			}
			log.Trace().
				Str("handler", "PollNetMapStream").
				Str("machine", m.Name).
				Str("channel", "pollData").
				Int("bytes", len(data)).
				Msg("Data from pollData channel written successfully")
			now := time.Now().UTC()
			m.LastSeen = &now
			m.LastSuccessfulUpdate = &now
			h.db.Save(&m)
			log.Trace().
				Str("handler", "PollNetMapStream").
				Str("machine", m.Name).
				Str("channel", "pollData").
				Int("bytes", len(data)).
				Msg("Machine updated successfully after sending pollData")
			return true

		case data := <-keepAliveChan:
			log.Trace().
				Str("handler", "PollNetMapStream").
				Str("machine", m.Name).
				Str("channel", "keepAlive").
				Int("bytes", len(data)).
				Msg("Sending keep alive message")
			_, err := w.Write(data)
			if err != nil {
				log.Error().
					Str("handler", "PollNetMapStream").
					Str("machine", m.Name).
					Str("channel", "keepAlive").
					Err(err).
					Msg("Cannot write keep alive message")
			}
			log.Trace().
				Str("handler", "PollNetMapStream").
				Str("machine", m.Name).
				Str("channel", "keepAlive").
				Int("bytes", len(data)).
				Msg("Keep alive sent successfully")
			now := time.Now().UTC()
			m.LastSeen = &now
			h.db.Save(&m)
			log.Trace().
				Str("handler", "PollNetMapStream").
				Str("machine", m.Name).
				Str("channel", "keepAlive").
				Int("bytes", len(data)).
				Msg("Machine updated successfully after sending keep alive")
			return true

		case <-updateChan:
			if h.isOutdated(&m) {
				log.Trace().
					Str("handler", "PollNetMapStream").
					Str("machine", m.Name).
					Str("channel", "update").
					Msg("Received a request for update")
				data, err := h.getMapResponse(mKey, req, m)
				if err != nil {
					log.Error().
						Str("handler", "PollNetMapStream").
						Str("machine", m.Name).
						Str("channel", "update").
						Err(err).
						Msg("Could not get the map update")
				}
				_, err = w.Write(*data)
				if err != nil {
					log.Error().
						Str("handler", "PollNetMapStream").
						Str("machine", m.Name).
						Str("channel", "update").
						Err(err).
						Msg("Could not write the map response")
				}
				log.Trace().
					Str("handler", "PollNetMapStream").
					Str("machine", m.Name).
					Str("channel", "update").
					Msg("Updated Map has been sent")

				// Keep track of the last successful update,
				// we sometimes end in a state were the update
				// is not picked up by a client and we use this
				// to determine if we should "force" an update.
				now := time.Now().UTC()
				m.LastSuccessfulUpdate = &now
				h.db.Save(&m)
			}
			return true

		case <-c.Request.Context().Done():
			log.Info().
				Str("handler", "PollNetMapStream").
				Str("machine", m.Name).
				Msg("The client has closed the connection")
			now := time.Now().UTC()
			m.LastSeen = &now
			h.db.Save(&m)

			cancelKeepAlive <- struct{}{}

			h.clientsUpdateChannels.Delete(m.ID)
			// close(updateChan)

			close(pollDataChan)

			close(keepAliveChan)

			return false
		}
	})
}

// TODO: Rename this function to schedule ...
func (h *Headscale) keepAlive(
	cancelChan <-chan struct{},
	keepAliveChan chan<- []byte,
	mKey wgkey.Key,
	req tailcfg.MapRequest,
	m Machine,
) {
	keepAliveTicker := time.NewTicker(60 * time.Second)
	updateCheckerTicker := time.NewTicker(30 * time.Second)

	for {
		select {
		case <-cancelChan:
			return

		case <-keepAliveTicker.C:
			data, err := h.getMapKeepAliveResponse(mKey, req, m)
			if err != nil {
				log.Error().
					Str("func", "keepAlive").
					Err(err).
					Msg("Error generating the keep alive msg")
				return
			}

			log.Debug().
				Str("func", "keepAlive").
				Str("machine", m.Name).
				Msg("Sending keepalive")
			keepAliveChan <- *data

		case <-updateCheckerTicker.C:
			err := h.UpdateMachine(&m)
			if err != nil {
				log.Error().
					Str("func", "keepAlive").
					Str("machine", m.Name).
					Err(err).
					Msg("Could not refresh machine details from database")
				return
			}
			if h.isOutdated(&m) {
				log.Debug().
					Str("func", "keepAlive").
					Str("machine", m.Name).
					Time("last_successful_update", *m.LastSuccessfulUpdate).
					Time("last_state_change", h.getLastStateChange()).
					Msgf("There has been updates since the last successful update to %s", m.Name)

				// TODO Error checking
				n, _ := m.toNode()
				h.requestUpdate(n)
			} else {
				log.Trace().
					Str("func", "keepAlive").
					Str("machine", m.Name).
					Time("last_successful_update", *m.LastSuccessfulUpdate).
					Time("last_state_change", h.getLastStateChange()).
					Msgf("%s is up to date", m.Name)
			}
		}
	}
}
