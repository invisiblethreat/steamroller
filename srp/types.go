package srp

import (
	"fmt"
	"strings"
	"sync"
)

// 2^16
const maxPort = 65536

// AllTargets holds the exploded arguments which are used for the Cartesian
// product to generate the set of atomic SingleTargets
type AllTargets struct {
	Addrs []string
	Port  int
	Proto string
}

// Load builds out atomic targets
func (at *AllTargets) Load(output chan SingleTarget, wg *sync.WaitGroup) {
	for _, addr := range at.Addrs {
		output <- SingleTarget{Addr: addr, Port: at.Port, Proto: at.Proto}
		wg.Add(1)
	}

}

// SingleTarget is an atomic entity to attempt a connection
type SingleTarget struct {
	Addr  string `json:"addr"`
	Port  int    `json:"port"`
	Proto string `json:"proto"`
}

// SteamRemotePlay holds the parsed response from the service.
type SteamRemotePlay struct {
	ExternalAddr  string   `json:"external_addr"`
	Amplification float64  `json:"amplification"`
	Name          string   `json:"name"`
	OS            string   `json:"os"`
	MACs          []string `json:"macs"`
	Addrs         []string `json:"addrs"`
}

// String, oddly, won't function as expected on a pointer, only by copy
func (s SteamRemotePlay) String() string {
	return fmt.Sprintf("                 Name: %s\n", s.Name) +
		fmt.Sprintf("        MAC addresses: %s\n", strings.Join(s.MACs, ", ")) +
		fmt.Sprintf("         IP addresses: %s\n", strings.Join(s.Addrs, ", ")) +
		fmt.Sprintf(" Amplification Factor: %f\n", s.Amplification)
}
