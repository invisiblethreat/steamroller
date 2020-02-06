package main

import (
	"time"

	"github.com/invisiblethreat/steamroller/srp"
)

type Result struct {
	Addr            string               `json:"addr"`
	Port            int                  `json:"port"`
	Proto           string               `json:"proto"`
	TimeStamp       time.Time            `json:"timestamp"`
	Plugin          string               `json:"plugin"`
	SteamRemotePlay *srp.SteamRemotePlay `json:"steam_remote_play,omitempty"`
	Unknown         string               `json:"unknown,omitempty"`
}

var modes = []string{
	"srp",
}

type Mode interface {
	Run()
}
