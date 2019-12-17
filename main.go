package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	ea "github.com/invisiblethreat/expandaddr"
	"github.com/invisiblethreat/steamroller/srp"
)

func main() {
	input := make(chan ea.SingleTarget)
	srpOutput := make(chan srp.SteamRemotePlay)
	unknown := make(chan srp.Unknown)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go srp.Worker(input, srpOutput, unknown, &wg)
	}
	go ResultHandler(srpOutput, &wg)
	go UnknownHandler(unknown, &wg)

	addrs, err := ea.ExpandAddrs(os.Args[1])
	if err != nil {
		fmt.Printf("%s", err.Error())
	}

	targets := ea.AllTargets{Addrs: addrs, Ports: []int{srp.PluginPort}, Protos: []string{srp.PluginProto}}
	targets.Load(input, &wg)
	close(input)
	wg.Wait()
	close(srpOutput)
}

// ResultHandler collects scan results and processes them. It is also the
// blocking function to ensure that results are not orpahaned
func ResultHandler(srpInput <-chan srp.SteamRemotePlay, wg *sync.WaitGroup) {
	for result := range srpInput {
		buf, err := json.Marshal(srpToResult(result))
		if err != nil {
			continue
		}
		fmt.Println(string(buf))
		// split the handling of the WaitGroup so that results can't be orphaned
		wg.Done()
	}
}
func UnknownHandler(input <-chan srp.Unknown, wg *sync.WaitGroup) {
	for unknown := range input {
		buf, err := json.Marshal(unknown)
		if err != nil {
			continue
		}
		fmt.Println(string(buf))
		wg.Done()
	}

}

func srpToResult(s srp.SteamRemotePlay) Result {
	return Result{
		Addr:            s.Target,
		TimeStamp:       time.Now().UTC(),
		Proto:           srp.PluginProto,
		Port:            srp.PluginPort,
		SteamRemotePlay: &s,
		Plugin:          srp.PluginString,
	}
}
