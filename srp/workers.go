package srp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"
)

// Worker checks the disposition of a host:port:protocol tuple and sends the
// result to the ResultHandler
func Worker(input chan SingleTarget, output chan SteamRemotePlay, wg *sync.WaitGroup) {
	for target := range input {
		raddr := net.UDPAddr{IP: net.ParseIP(target.Addr),
			Port: SrpPort}
		conn, err := net.DialUDP(SrpProto, nil, &raddr)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			wg.Done()
			continue
		}
		defer conn.Close()
		err = conn.SetReadDeadline(time.Now().Add(1 * time.Second))
		res := make([]byte, 1024)

		conn.Write(steamHello)
		_, err = bufio.NewReader(conn).Read(res)
		conn.Close()

		if err, ok := err.(*net.OpError); ok && err.Timeout() {
			wg.Done()
			continue
		}

		//protect the parser- it's fragile
		isSrp := true
		for i := 0; i < len(serverHello); i++ {

			if res[i] != serverHello[i] {
				isSrp = false
				break
			}
		}
		if !isSrp {
			wg.Done()

			continue
		}
		//fmt.Printf("address: %s", target.Addr)
		//fmt.Printf("%x\n", res)
		parsed := parse(res)
		parsed.ExternalAddr = target.Addr
		output <- parsed
	}
}

// ResultHandler collects scan results and processes them. It is also the
// blocking function to ensure that results are not orpahaned
func ResultHandler(input <-chan SteamRemotePlay, wg *sync.WaitGroup) {
	for result := range input {
		buf, err := json.Marshal(result)
		if err != nil {
			continue
		}
		fmt.Println(string(buf))

		wg.Done()
	}
}
