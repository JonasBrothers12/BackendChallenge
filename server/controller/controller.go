package controller

import (
	"github.com/JonasBrothers12/BackendChallenge/server/controller/user"
	"github.com/JonasBrothers12/BackendChallenge/service"
	"github.com/JonasBrothers12/BackendChallenge/util"
	"github.com/rs/zerolog"
)

type ControllerManager struct{
	UserController *user.ControllerUser
}

func NewController(svc *service.Service, logger *zerolog.Logger) *ControllerManager {
	resutil := util.New(logger)

	return &ControllerManager{
		UserController:   user.NewController(svc, resutil),
	}
}