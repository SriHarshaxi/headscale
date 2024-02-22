package headscale

import (
	"encoding/json"
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
	errorEmptyPolicy        = Error("empty policy")
	errorInvalidAction      = Error("invalid action")
	errorInvalidUserSection = Error("invalid user section")
	errorInvalidGroup       = Error("invalid group")
	errorInvalidTag         = Error("invalid tag")
	errorInvalidNamespace   = Error("invalid namespace")
	errorInvalidPortFormat  = Error("invalid port format")
)

const (
	PORT_RANGE_BEGIN     = 0
	PORT_RANGE_END       = 65535
	BASE_10              = 10
	BIT_SIZE_16          = 16
	EXPECTED_TOKEN_ITEMS = 2
)

// LoadACLPolicy loads the ACL policy from the specify path, and generates the ACL rules.
func (h *Headscale) LoadACLPolicy(path string) error {
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
		return errorEmptyPolicy
	}

	h.aclPolicy = &policy
	rules, err := h.generateACLRules()
	if err != nil {
		return err
	}
	h.aclRules = rules

	return nil
}

func (h *Headscale) generateACLRules() ([]tailcfg.FilterRule, error) {
	rules := []tailcfg.FilterRule{}

	for index, acl := range h.aclPolicy.ACLs {
		if acl.Action != "accept" {
			return nil, errorInvalidAction
		}

		filterRule := tailcfg.FilterRule{}

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
		filterRule.SrcIPs = srcIPs

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
	if len(tokens) < EXPECTED_TOKEN_ITEMS || len(tokens) > 3 {
		return nil, errorInvalidPortFormat
	}

	var alias string
	// We can have here stuff like:
	// git-server:*
	// 192.168.1.0/24:22
	// tag:montreal-webserver:80,443
	// tag:api-server:443
	// example-host-1:*
	if len(tokens) == EXPECTED_TOKEN_ITEMS {
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

func (h *Headscale) expandAlias(alias string) ([]string, error) {
	if alias == "*" {
		return []string{"*"}, nil
	}

	if strings.HasPrefix(alias, "group:") {
		if _, ok := h.aclPolicy.Groups[alias]; !ok {
			return nil, errorInvalidGroup
		}
		ips := []string{}
		for _, n := range h.aclPolicy.Groups[alias] {
			nodes, err := h.ListMachinesInNamespace(n)
			if err != nil {
				return nil, errorInvalidNamespace
			}
			for _, node := range nodes {
				ips = append(ips, node.IPAddress)
			}
		}

		return ips, nil
	}

	if strings.HasPrefix(alias, "tag:") {
		if _, ok := h.aclPolicy.TagOwners[alias]; !ok {
			return nil, errorInvalidTag
		}

		// This will have HORRIBLE performance.
		// We need to change the data model to better store tags
		machines := []Machine{}
		if err := h.db.Where("registered").Find(&machines).Error; err != nil {
			return nil, err
		}
		ips := []string{}
		for _, machine := range machines {
			hostinfo := tailcfg.Hostinfo{}
			if len(machine.HostInfo) != 0 {
				hi, err := machine.HostInfo.MarshalJSON()
				if err != nil {
					return nil, err
				}
				err = json.Unmarshal(hi, &hostinfo)
				if err != nil {
					return nil, err
				}

				// FIXME: Check TagOwners allows this
				for _, t := range hostinfo.RequestTags {
					if alias[4:] == t {
						ips = append(ips, machine.IPAddress)

						break
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
			ips = append(ips, n.IPAddress)
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

	return nil, errorInvalidUserSection
}

func (h *Headscale) expandPorts(portsStr string) (*[]tailcfg.PortRange, error) {
	if portsStr == "*" {
		return &[]tailcfg.PortRange{
			{First: PORT_RANGE_BEGIN, Last: PORT_RANGE_END},
		}, nil
	}

	ports := []tailcfg.PortRange{}
	for _, portStr := range strings.Split(portsStr, ",") {
		rang := strings.Split(portStr, "-")
		switch len(rang) {
		case 1:
			port, err := strconv.ParseUint(rang[0], BASE_10, BIT_SIZE_16)
			if err != nil {
				return nil, err
			}
			ports = append(ports, tailcfg.PortRange{
				First: uint16(port),
				Last:  uint16(port),
			})

		case EXPECTED_TOKEN_ITEMS:
			start, err := strconv.ParseUint(rang[0], BASE_10, BIT_SIZE_16)
			if err != nil {
				return nil, err
			}
			last, err := strconv.ParseUint(rang[1], BASE_10, BIT_SIZE_16)
			if err != nil {
				return nil, err
			}
			ports = append(ports, tailcfg.PortRange{
				First: uint16(start),
				Last:  uint16(last),
			})

		default:
			return nil, errorInvalidPortFormat
		}
	}

	return &ports, nil
}
