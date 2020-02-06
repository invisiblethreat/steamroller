package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/TV4/env"
	ea "github.com/invisiblethreat/expandaddr"
	"github.com/invisiblethreat/steamroller/srp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

type cliOptions struct {
	debug    bool
	pprof    bool
	unknowns bool
	mode     string
	receiver string
	target   string
	file     string
	workers  int
}

func main() {

	var options cliOptions
	pflag.BoolVarP(&options.debug, "debug", "d", false, "Enable debugging")
	pflag.BoolVarP(&options.pprof, "pprof", "p", false, "Enable pprof")
	pflag.BoolVarP(&options.unknowns, "unknowns", "u", false, "Output unknown results for furhter analysis")
	pflag.StringVarP(&options.mode, "mode", "m", "", "Processing mode: valid options: srp")
	pflag.StringVarP(&options.receiver, "receiver", "r", "", "address:port to receive responses")
	pflag.StringVarP(&options.target, "target", "t", "", "Target: single address, or CIDR block")
	pflag.StringVarP(&options.file, "file", "f", "", "File of IP addresses to probe")
	pflag.IntVarP(&options.workers, "workers", "w", 100, "Number of workers to deploy. Warning: you can DoS yourself!")
	pflag.Parse()

	logging(options.debug)
	pprof(options.pprof)

	input := make(chan ea.SingleTarget)

	wg := sync.WaitGroup{}

	if options.mode == "srp" {
		config := options.toSrpConfig()
		config.Input = input
		config.Wg = &wg
		go srp.Run(config)
	} else {
		panic("No valid mode selected. Use -h to see modes")
	}

	if options.target != "" {
		addrs, err := ea.ExpandAddrs(options.target)
		if err != nil {
			fmt.Printf("%s", err.Error())
		}

		targets := ea.AllTargets{Addrs: addrs, Ports: []int{srp.PluginPort}, Protos: []string{srp.PluginProto}}
		targets.Load(input, &wg)
	}

	if options.file != "" {
		addrs, err := loadFile(options.file)
		if err != nil {
			panic(fmt.Sprintf("Error loading file %s", err.Error()))
		}
		targets := ea.AllTargets{Addrs: addrs, Ports: []int{srp.PluginPort}, Protos: []string{srp.PluginProto}}
		targets.Load(input, &wg)
	}
	close(input)
	wg.Wait()
}

func logging(debug bool) {
	if debug || (strings.ToLower(env.String("LOG_DEBUG", "")) == "true") {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug("Running in debug mode.")
	}
	logrus.WithField("level", logrus.GetLevel().String()).Info("Logging level")
}

func pprof(enabled bool) {
	if enabled {
		go func() {
			err := http.ListenAndServe("localhost:6060", nil)
			if err != nil {
				logrus.WithError(err).Errorf("Error starting pprof")
			}
		}()
	}
}

func (c cliOptions) toSrpConfig() srp.SrpConfig {
	return srp.SrpConfig{
		Workers:  c.workers,
		Unknown:  c.unknowns,
		Receiver: c.receiver,
	}
}
