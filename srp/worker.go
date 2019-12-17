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
func Worker(input chan expandaddr.SingleTarget, output chan SteamRemotePlay, unknown chan []byte, wg *sync.WaitGroup) {
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

		if !bytes.Equal(serverHello, res[0:len(serverHello)]) {
			unknown <- res
			continue
		}
		parsed, _ := parse(res)

		output <- parsed
	}
}

/*
func buildSrpResult(t SingleTarget, s SteamRemotePlay) Result {
	return Result{Addr: t.Addr, Proto: t.Proto, TimeStamp: time.Now().UTC(),
		SteamRemotePlay: &s, Plugin: PluginString}
}
func buildUnknownResult(t SingleTarget, b []byte) Result {
	return Result{Addr: t.Addr, Proto: t.Proto, TimeStamp: time.Now().UTC(),
		Unknown: fmt.Sprintf("%x", b), Plugin: PluginString}
}
*/
