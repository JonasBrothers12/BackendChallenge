package main

import (
	"os"

	"github.com/JonasBrothers12/BackendChallenge/config"
	"github.com/JonasBrothers12/BackendChallenge/database"
	"github.com/JonasBrothers12/BackendChallenge/logger"
	"github.com/JonasBrothers12/BackendChallenge/server"
	"github.com/JonasBrothers12/BackendChallenge/server/controller"
	"github.com/JonasBrothers12/BackendChallenge/service"
)

func main() {
	cfg := config.Getconfigs()
	logger := logger.Initiallogger(cfg.InternalConfig)
	repo, err := database.NewDatabaseAction(cfg)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize DatabaseManager")
		os.Exit(1)
	}
	svc := service.NewService(cfg,logger,repo)
	ctrl := controller.NewController(svc,logger)
	svr := server.NewServer(cfg,logger,ctrl)
	svr.InitializingServer()

}
