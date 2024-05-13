package server

import (
	"github.com/JonasBrothers12/BackendChallenge/config"
	middleware "github.com/JonasBrothers12/BackendChallenge/server/Middleware"
	"github.com/JonasBrothers12/BackendChallenge/server/controller"
	"github.com/JonasBrothers12/BackendChallenge/server/routes"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Server struct {
	cfg    *config.Config
	svr    *gin.Engine
	logger *zerolog.Logger
	ctrl   *controller.ControllerManager
}


func (s *Server)InitializingServer() error{
  s.logger.Info().Msg("starting server")  
  if err := s.svr.Run(":8080"); err != nil{
	return err
  }
  return nil
}


func NewServer(cfg *config.Config,logger *zerolog.Logger,ctrl *controller.ControllerManager) *Server{
	svr := gin.New()
	middleware.SetMiddlewares(svr,cfg)
	router.RegisterRoutes(svr,ctrl)
	return &Server{
		cfg:    cfg,
		svr:    svr,
		logger: logger,
		ctrl:   ctrl,
	}
}
