package headscale

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/tailscale/hujson"
	"inet.af/netaddr"
	"tailscale.com/tailcfg"
)

const (
	errEmptyPolicy        = Error("empty policy")
	errInvalidAction      = Error("invalid action")
	errInvalidUserSection = Error("invalid user section")
	errInvalidGroup       = Error("invalid group")
	errInvalidTag         = Error("invalid tag")
	errInvalidNamespace   = Error("invalid namespace")
	errInvalidPortFormat  = Error("invalid port format")
)

const (
	Base8              = 8
	Base10             = 10
	BitSize16          = 16
	BitSize32          = 32
	BitSize64          = 64
	portRangeBegin     = 0
	portRangeEnd       = 65535
	expectedTokenItems = 2
)

// LoadACLPolicy loads the ACL policy from the specify path, and generates the ACL rules.
func (h *Headscale) LoadACLPolicy(path string) error {
	log.Debug().
		Str("func", "LoadACLPolicy").
		Str("path", path).
		Msg("Loading ACL policy from path")

	policyFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer policyFile.Close()

	var policy ACLPolicy
	policyBytes, err := io.ReadAll(policyFile)
	if err != nil {
		return err
	}

	ast, err := hujson.Parse(policyBytes)
	if err != nil {
		return err
	}
	ast.Standardize()
	policyBytes = ast.Pack()
	err = json.Unmarshal(policyBytes, &policy)
	if err != nil {
		return err
	}
	if policy.IsZero() {
		return errEmptyPolicy
	}

	h.aclPolicy = &policy
	return h.UpdateACLRules()
}

func (h *Headscale) UpdateACLRules() error {
	rules, err := h.generateACLRules()
	if err != nil {
		return err
	}
	log.Trace().Interface("ACL", rules).Msg("ACL rules generated")
	h.aclRules = rules
	return nil
}

func (h *Headscale) generateACLRules() ([]tailcfg.FilterRule, error) {
	rules := []tailcfg.FilterRule{}

	for index, acl := range h.aclPolicy.ACLs {
		if acl.Action != "accept" {
			return nil, errInvalidAction
		}

		srcIPs := []string{}
		for innerIndex, user := range acl.Users {
			srcs, err := h.generateACLPolicySrcIP(user)
			if err != nil {
				log.Error().
					Msgf("Error parsing ACL %d, User %d", index, innerIndex)

				return nil, err
			}
			srcIPs = append(srcIPs, srcs...)
		}

		destPorts := []tailcfg.NetPortRange{}
		for innerIndex, ports := range acl.Ports {
			dests, err := h.generateACLPolicyDestPorts(ports)
			if err != nil {
				log.Error().
					Msgf("Error parsing ACL %d, Port %d", index, innerIndex)

				return nil, err
			}
			destPorts = append(destPorts, dests...)
		}

		rules = append(rules, tailcfg.FilterRule{
			SrcIPs:   srcIPs,
			DstPorts: destPorts,
		})
	}

	return rules, nil
}

func (h *Headscale) generateACLPolicySrcIP(u string) ([]string, error) {
	return h.expandAlias(u)
}

func (h *Headscale) generateACLPolicyDestPorts(
	d string,
) ([]tailcfg.NetPortRange, error) {
	tokens := strings.Split(d, ":")
	if len(tokens) < expectedTokenItems || len(tokens) > 3 {
		return nil, errInvalidPortFormat
	}

	var alias string
	// We can have here stuff like:
	// git-server:*
	// 192.168.1.0/24:22
	// tag:montreal-webserver:80,443
	// tag:api-server:443
	// example-host-1:*
	if len(tokens) == expectedTokenItems {
		alias = tokens[0]
	} else {
		alias = fmt.Sprintf("%s:%s", tokens[0], tokens[1])
	}

	expanded, err := h.expandAlias(alias)
	if err != nil {
		return nil, err
	}
	ports, err := h.expandPorts(tokens[len(tokens)-1])
	if err != nil {
		return nil, err
	}

	dests := []tailcfg.NetPortRange{}
	for _, d := range expanded {
		for _, p := range *ports {
			pr := tailcfg.NetPortRange{
				IP:    d,
				Ports: p,
			}
			dests = append(dests, pr)
		}
	}

	return dests, nil
}

// expandalias has an input of either
// - a namespace
// - a group
// - a tag
// and transform these in IPAddresses
func (h *Headscale) expandAlias(alias string) ([]string, error) {
	if alias == "*" {
		return []string{"*"}, nil
	}

	if strings.HasPrefix(alias, "group:") {
		namespaces, err := h.expandGroup(alias)
		if err != nil {
			return nil, err
		}
		ips := []string{}
		for _, n := range namespaces {
			nodes, err := h.ListMachinesInNamespace(n)
			if err != nil {
				return nil, errInvalidNamespace
			}
			for _, node := range nodes {
				ips = append(ips, node.IPAddresses.ToStringSlice()...)
			}
		}

		return ips, nil
	}

	if strings.HasPrefix(alias, "tag:") {
		var ips []string
		owners, err := h.expandTagOwners(alias)
		if err != nil {
			return nil, err
		}
		for _, namespace := range owners {
			machines, err := h.ListMachinesInNamespace(namespace)
			if err != nil {
				if errors.Is(err, errNamespaceNotFound) {
					continue
				} else {
					return nil, err
				}
			}
			for _, machine := range machines {
				if len(machine.HostInfo) == 0 {
					continue
				}
				hi, err := machine.GetHostInfo()
				if err != nil {
					return nil, err
				}
				for _, t := range hi.RequestTags {
					if alias == t {
						ips = append(ips, machine.IPAddresses.ToStringSlice()...)
					}
				}
			}
		}
		return ips, nil
	}

	n, err := h.GetNamespace(alias)
	if err == nil {
		nodes, err := h.ListMachinesInNamespace(n.Name)
		if err != nil {
			return nil, err
		}
		ips := []string{}
		for _, n := range nodes {
			ips = append(ips, n.IPAddresses.ToStringSlice()...)
		}

		return ips, nil
	}

	if h, ok := h.aclPolicy.Hosts[alias]; ok {
		return []string{h.String()}, nil
	}

	ip, err := netaddr.ParseIP(alias)
	if err == nil {
		return []string{ip.String()}, nil
	}

	cidr, err := netaddr.ParseIPPrefix(alias)
	if err == nil {
		return []string{cidr.String()}, nil
	}

	return nil, errInvalidUserSection
}

// expandTagOwners will return a list of namespace. An owner can be either a namespace or a group
// a group cannot be composed of groups
func (h *Headscale) expandTagOwners(owner string) ([]string, error) {
	var owners []string
	ows, ok := h.aclPolicy.TagOwners[owner]
	if !ok {
		return []string{}, fmt.Errorf("%w. %v isn't owned by a TagOwner. Please add one first. https://tailscale.com/kb/1018/acls/#tag-owners", errInvalidTag, owner)
	}
	for _, ow := range ows {
		if strings.HasPrefix(ow, "group:") {
			gs, err := h.expandGroup(ow)
			if err != nil {
				return []string{}, err
			}
			owners = append(owners, gs...)
		} else {
			owners = append(owners, ow)
		}
	}
	return owners, nil
}

// expandGroup will return the list of namespace inside the group
// after some validation
func (h *Headscale) expandGroup(group string) ([]string, error) {
	gs, ok := h.aclPolicy.Groups[group]
	if !ok {
		return []string{}, fmt.Errorf("group %v isn't registered. %w", group, errInvalidGroup)
	}
	for _, g := range gs {
		if strings.HasPrefix(g, "group:") {
			return []string{}, fmt.Errorf("%w. A group cannot be composed of groups. https://tailscale.com/kb/1018/acls/#groups", errInvalidGroup)
		}
	}
	return gs, nil
}

func (h *Headscale) expandPorts(portsStr string) (*[]tailcfg.PortRange, error) {
	if portsStr == "*" {
		return &[]tailcfg.PortRange{
			{First: portRangeBegin, Last: portRangeEnd},
		}, nil
	}

	ports := []tailcfg.PortRange{}
	for _, portStr := range strings.Split(portsStr, ",") {
		rang := strings.Split(portStr, "-")
		switch len(rang) {
		case 1:
			port, err := strconv.ParseUint(rang[0], Base10, BitSize16)
			if err != nil {
				return nil, err
			}
			ports = append(ports, tailcfg.PortRange{
				First: uint16(port),
				Last:  uint16(port),
			})

		case expectedTokenItems:
			start, err := strconv.ParseUint(rang[0], Base10, BitSize16)
			if err != nil {
				return nil, err
			}
			last, err := strconv.ParseUint(rang[1], Base10, BitSize16)
			if err != nil {
				return nil, err
			}
			ports = append(ports, tailcfg.PortRange{
				First: uint16(start),
				Last:  uint16(last),
			})

		default:
			return nil, errInvalidPortFormat
		}
	}

	return &ports, nil
}
