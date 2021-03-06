package srp

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/invisiblethreat/expandaddr"
)

// SteamRemotePlay holds the parsed response from the service.
type SteamRemotePlay struct {
	Target        string    `json:"target"`
	Port          int       `json:"port"`
	Plugin        string    `json:"plugin"`
	Proto         string    `json:"proto"`
	Timestamp     time.Time `json:"timestamp"`
	Amplification float64   `json:"amplification"`
	Version       int       `json:"version"`
	Name          string    `json:"name"`
	OS            string    `json:"os"`
	MACs          []string  `json:"macs"`
	Addrs         []string  `json:"addrs"`
}

type Unknown struct {
	Target    string    `json:"target"`
	Port      int       `json:"port"`
	Timestamp time.Time `json:"timestamp"`
	Plugin    string    `json:"plugin"`
	Proto     string    `json:"proto"`
	Payload   string    `json:"payload"`
}

type SrpConfig struct {
	Workers  int
	Unknown  bool
	Input    chan expandaddr.SingleTarget
	Receiver string
	Wg       *sync.WaitGroup
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
