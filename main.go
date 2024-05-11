package main

import (
	"github.com/JonasBrothers12/BackendChallenge/config"
	"github.com/JonasBrothers12/BackendChallenge/server"
)

func main() {
	cfg := config.Getconfigs()
	server.InitializingServer(cfg)
}