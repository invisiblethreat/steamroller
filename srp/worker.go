package srp

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/invisiblethreat/expandaddr"
	"github.com/invisiblethreat/locate"
	"github.com/sirupsen/logrus"
)

func Run(c SrpConfig) {
	log := logrus.WithField("loc", locate.WhereAmI())
	results := make(chan SteamRemotePlay)
	unknown := make(chan Unknown)
	log.Debug("Starting SRP handlers")
	go ResultHandler(results, c.Wg)
	go UnknownHandler(unknown, c.Wg)

	log.Debugf("Starting %d SRP workers", c.Workers)
	for i := 0; i < c.Workers; i++ {
		log.Debugf("Started worker %d", i+1)
		go Worker(c.Input, results, unknown, c.Unknown, nil, c.Wg)
	}
}

// Worker checks the disposition of a host:port:protocol tuple and sends the
// result to the ResultHandler
func Worker(input <-chan expandaddr.SingleTarget, output chan<- SteamRemotePlay, unknown chan<- Unknown, keepUnknown bool, receiver *net.UDPAddr, wg *sync.WaitGroup) {
	log := logrus.WithField("loc", locate.WhereAmI())
	for target := range input {
		raddr := net.UDPAddr{IP: net.ParseIP(target.Addr),
			Port: target.Port}
		conn, err := net.DialUDP(PluginProto, receiver, &raddr)
		if err != nil {
			log.WithError(err).Errorf("Error dialing")
			wg.Done()
			continue
		}

		if receiver != nil {
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
			log.Debugf("%s timed out", target.Addr)
			wg.Done()
			continue
		}

		remoteHello := res[0:len(serverHello)]
		if !bytes.Equal(serverHello, remoteHello) {

			if bytes.Equal(emptyResponse, res[0:len(emptyResponse)]) {
				wg.Done()
				continue
			} else {
				if keepUnknown {
					unknown <- Unknown{
						Target:  target.Addr,
						Port:    target.Port,
						Plugin:  PluginString,
						Proto:   target.Proto,
						Payload: fmt.Sprintf("%x", res)}
					continue
				} else {
					wg.Done()
					continue
				}
			}
		}
		logrus.Debugf("Parsing results for %s", target.Addr)
		parsed, err := parse(res)
		if err != nil {
			logrus.WithError(err).Error("Error parsing results")
		}
		parsed.Target = target.Addr
		parsed.Timestamp = time.Now().UTC()
		parsed.Proto = PluginProto
		parsed.Port = PluginPort
		parsed.Plugin = PluginString
		output <- parsed
	}
}

func ResultHandler(srpInput <-chan SteamRemotePlay, wg *sync.WaitGroup) {
	for result := range srpInput {
		log := logrus.WithField("loc", locate.WhereAmI())
		log.Debugf("Got result for %s", result.Target)
		buf, err := json.Marshal(result)
		if err != nil {
			log.WithError(err).Error("Marshal failed")
			wg.Done()
			continue
		}
		fmt.Println(string(buf))
		wg.Done()
	}
}

func UnknownHandler(input <-chan Unknown, wg *sync.WaitGroup) {
	log := logrus.WithField("loc", locate.WhereAmI())
	for unknown := range input {
		buf, err := json.Marshal(unknown)
		if err != nil {
			log.WithError(err).Error("Marshal failed")
			wg.Done()
			continue
		}
		fmt.Println(string(buf))
		wg.Done()
	}

}
