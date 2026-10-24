package runkey

import (
	"fmt"
	"net"
	"net/http"
	"net/netip"
	"strings"
)

// Callers tells a call from the job's own machine (the agent calling the API
// back) from any other. Taking run keys from Callers only makes a key that
// leaks useless elsewhere.
type Callers struct {
	networks []netip.Prefix
}

// Loopback: calls from this machine only, the production case.
func Loopback() Callers {
	return Callers{networks: []netip.Prefix{netip.MustParsePrefix("127.0.0.0/8"), netip.MustParsePrefix("::1/128")}}
}

// ParseCallers takes comma-separated CIDRs ("172.16.0.0/12,192.168.0.0/16"
// for an API in Docker in dev); empty: Loopback.
func ParseCallers(cidrs string) (Callers, error) {
	if strings.TrimSpace(cidrs) == "" {
		return Loopback(), nil
	}
	var networks []netip.Prefix
	for _, cidr := range strings.Split(cidrs, ",") {
		network, err := netip.ParsePrefix(strings.TrimSpace(cidr))
		if err != nil {
			return Callers{}, fmt.Errorf("runkey: ParseCallers: %w", err)
		}
		networks = append(networks, network)
	}
	return Callers{networks: networks}, nil
}

// Allow: the request comes straight from one of the networks. A request
// through a reverse proxy, on the same machine too, carries X-Forwarded-For
// (Go's httputil.ReverseProxy always sets it): refused. A zero Callers allows
// nothing.
func (c Callers) Allow(r *http.Request) bool {
	if r.Header.Get("X-Forwarded-For") != "" {
		return false
	}
	addr, ok := remoteAddr(r)
	if !ok {
		return false
	}
	for _, network := range c.networks {
		if network.Contains(addr) {
			return true
		}
	}
	return false
}

// remoteAddr is the IP of the request's peer, IPv4 in IPv6 unmapped.
func remoteAddr(r *http.Request) (netip.Addr, bool) {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return netip.Addr{}, false
	}
	addr, err := netip.ParseAddr(host)
	if err != nil {
		return netip.Addr{}, false
	}
	return addr.Unmap(), true
}
