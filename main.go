package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/invisiblethreat/steamroller/srp"
)

func main() {
	input := make(chan srp.SingleTarget)
	output := make(chan srp.SteamRemotePlay)
	wg := sync.WaitGroup{}
	for i := 0; i < 256; i++ {
		go srp.Worker(input, output, &wg)
	}
	go srp.ResultHandler(output, &wg)

	addrs, err := srp.ExpandAddrs(os.Args[1])
	if err != nil {
		fmt.Printf("%s", err.Error())
	}

	targets := srp.AllTargets{Addrs: addrs, Port: srp.SrpPort, Proto: srp.SrpProto}
	targets.Load(input, &wg)
	close(input)
	wg.Wait()
	close(output)
}
