package service

import (
	"github.com/JonasBrothers12/BackendChallenge/config"
	"github.com/JonasBrothers12/BackendChallenge/database"
	"github.com/rs/zerolog"
)

type Service struct {
	User *UserService
}

func NewService(cfg *config.Config, logger *zerolog.Logger, repo *database.DatabaseManager) *Service {
	return &Service{
		User: NewUserService(cfg, logger, repo),
	}
}
