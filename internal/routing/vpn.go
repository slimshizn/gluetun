package routing

import (
	"errors"
	"fmt"
	"net/netip"

	"github.com/qdm12/gluetun/internal/netlink"
	"golang.org/x/sys/unix"
)

var (
	ErrVPNLocalGatewayIPNotFound       = errors.New("VPN local gateway IP address not found")
	ErrVPNLocalGatewayIPv6NotSupported = errors.New("VPN local gateway IPv6 address not supported")
)

func (r *Routing) VPNLocalGatewayIP(vpnIntf string) (ip netip.Addr, err error) {
	vpnLink, err := r.netLinker.LinkByName(vpnIntf)
	if err != nil {
		return ip, fmt.Errorf("finding link %s: %w", vpnIntf, err)
	}
	vpnLinkIndex := vpnLink.Index

	routes, err := r.netLinker.RouteList(netlink.FamilyAll)
	if err != nil {
		return ip, fmt.Errorf("listing routes: %w", err)
	}
	for _, route := range routes {
		if route.LinkIndex != vpnLinkIndex {
			continue
		}

		switch {
		case route.Dst.IsValid() && route.Dst.Addr().IsUnspecified() && route.Gw.IsValid(): // OpenVPN
			return route.Gw, nil
		case route.Dst.IsSingleIP() &&
			route.Dst.Addr().Compare(route.Src) == 0 &&
			route.Table == unix.RT_TABLE_LOCAL: // Wireguard
			route.Src = route.Src.Unmap()
			if route.Src.Is6() {
				return netip.Addr{}, fmt.Errorf("%w: %s", ErrVPNLocalGatewayIPv6NotSupported, route.Src)
			}
			bytes := route.Src.As4()
			// force last byte to 1 to get the VPN gateway IP
			// This is not necessarily bullet proof but it seems to work.
			bytes[3] = 1
			return netip.AddrFrom4(bytes), nil
		}
	}
	return ip, fmt.Errorf("%w: in %d routes", ErrVPNLocalGatewayIPNotFound, len(routes))
}
