package srp

import (
	"net"
	"strconv"
	"strings"
)

func ExpandAddrs(addrs string) ([]string, error) {
	// Discrete IP addresses, comma seperated.
	if strings.Contains(addrs, ",") {
		return expandCommaDelim(addrs), nil
		// CIDR that needs expanded
	} else if strings.Contains(addrs, "/") {
		return expandCIDR(addrs)
	}
	// Single address inserted into slice
	return []string{addrs}, nil

}

func expandCommaDelim(s string) []string {
	return strings.Split(s, ",")
}

func expandCIDR(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); addrInc(ip) {
		ips = append(ips, ip.String())
	}
	// remove network address and broadcast address
	if len(ips) == 1 {
		return ips, nil
	}
	return ips[1 : len(ips)-1], nil
}
func addrInc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func ExpandPorts(port string) ([]string, error) {
	var ports []string
	if strings.Contains(port, "-") {
		ranges := strings.Split(port, "-")
		start, err := strconv.Atoi(ranges[0])
		if err != nil {
			return nil, err
		}
		end, err := strconv.Atoi(ranges[1])
		if err != nil {
			return nil, err
		}
		if end > maxPort {
			end = maxPort
		}
		for i := start; i <= end; i++ {
			ports = append(ports, strconv.Itoa(i))
		}

	} else if strings.Contains(port, ",") {
		return expandCommaDelim(port), nil
	} else {
		ports = []string{port}
	}
	return ports, nil
}
