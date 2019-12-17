package srp

import (
	"fmt"
	"strings"
)

// SteamRemotePlay holds the parsed response from the service.
type SteamRemotePlay struct {
	Target        string   `json:"target"`
	Amplification float64  `json:"amplification"`
	Version       int      `json:"version"`
	Name          string   `json:"name"`
	OS            string   `json:"os"`
	MACs          []string `json:"macs"`
	Addrs         []string `json:"addrs"`
}

type Unknown struct {
	Target  string `json:"target"`
	Port    int    `json:"port"`
	Plugin  string `json:"plugin"`
	Proto   string `json:"proto"`
	Payload string `json:"payload"`
}

// String, oddly, won't function as expected on a pointer, only by copy
func (s SteamRemotePlay) String() string {
	return fmt.Sprintf("               Target: %s\n", s.Target) +
		fmt.Sprintf("                 Name: %s\n", s.Name) +
		fmt.Sprintf("              Version: %d\n", s.Version) +
		fmt.Sprintf("        MAC addresses: %s\n", strings.Join(s.MACs, ", ")) +
		fmt.Sprintf("         IP addresses: %s\n", strings.Join(s.Addrs, ", ")) +
		fmt.Sprintf("                   OS: %s\n", s.OS) +
		fmt.Sprintf(" Amplification Factor: %f\n", s.Amplification)
}
