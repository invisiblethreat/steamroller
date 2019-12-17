package srp

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/invisiblethreat/expandaddr"
)

// Worker checks the disposition of a host:port:protocol tuple and sends the
// result to the ResultHandler
func Worker(input <-chan expandaddr.SingleTarget, output chan<- SteamRemotePlay, unknown chan<- Unknown, wg *sync.WaitGroup) {
	for target := range input {
		raddr := net.UDPAddr{IP: net.ParseIP(target.Addr),
			Port: target.Port}
		conn, err := net.DialUDP(PluginProto, nil, &raddr)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			wg.Done()
			continue
		}
		defer conn.Close()
		err = conn.SetReadDeadline(time.Now().Add(1 * time.Second))
		res := make([]byte, 1024)

		conn.Write(clientHello)
		_, err = bufio.NewReader(conn).Read(res)
		conn.Close()

		if err, ok := err.(*net.OpError); ok && err.Timeout() {
			wg.Done()
			continue
		}

		remoteHello := res[0:len(serverHello)]
		if !bytes.Equal(serverHello, remoteHello) {
			if bytes.Equal(emptyResponse, res[0:len(emptyResponse)]) {
				wg.Done()
				continue
			} else {
				unknown <- Unknown{
					Target:  target.Addr,
					Port:    target.Port,
					Plugin:  PluginString,
					Proto:   target.Proto,
					Payload: fmt.Sprintf("%x", res)}
				continue
			}
		}
		parsed, _ := parse(res)
		parsed.Target = target.Addr
		output <- parsed
	}
}
