package updater

import (
	"context"
	"fmt"
	"sort"

	"github.com/qdm12/gluetun/internal/constants/vpn"
	"github.com/qdm12/gluetun/internal/models"
	"github.com/qdm12/gluetun/internal/provider/common"
)

func (u *Updater) FetchServers(ctx context.Context, minServers int) (
	servers []models.Server, err error,
) {
	tcpHostToURL, udpHostToURL, err := getAllHostToURL(ctx, u.client)
	if err != nil {
		return nil, err
	}

	hosts := getUniqueHosts(tcpHostToURL, udpHostToURL)

	if len(hosts) < minServers {
		return nil, fmt.Errorf("%w: %d and expected at least %d",
			common.ErrNotEnoughServers, len(hosts), minServers)
	}

	resolveSettings := parallelResolverSettings(hosts)
	hostToIPs, warnings, err := u.parallelResolver.Resolve(ctx, resolveSettings)
	for _, warning := range warnings {
		u.warner.Warn(warning)
	}
	if err != nil {
		return nil, err
	}

	if len(hostToIPs) < minServers {
		return nil, fmt.Errorf("%w: %d and expected at least %d",
			common.ErrNotEnoughServers, len(servers), minServers)
	}

	servers = make([]models.Server, 0, len(hostToIPs))
	for host, IPs := range hostToIPs {
		tcpURL, tcp := tcpHostToURL[host]
		udpURL, udp := udpHostToURL[host]

		// These two are only used to extract the country, region and city.
		var url, protocol string
		if tcp {
			url = tcpURL
			protocol = "TCP"
		} else if udp {
			url = udpURL
			protocol = "UDP"
		}
		country, region, city := parseOpenvpnURL(url, protocol)

		server := models.Server{
			VPN:      vpn.OpenVPN,
			Country:  country,
			Region:   region,
			City:     city,
			Hostname: host,
			IPs:      IPs,
			TCP:      tcp,
			UDP:      udp,
		}
		servers = append(servers, server)
	}

	sort.Sort(models.SortableServers(servers))

	return servers, nil
}
