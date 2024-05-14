package service

import (
	"github.com/JonasBrothers12/BackendChallenge/config"
	"github.com/JonasBrothers12/BackendChallenge/database"
	"github.com/rs/zerolog"
)

type Service struct {
	User 	*UserService
	Wallet  *WalletService
}

func NewService(cfg *config.Config, logger *zerolog.Logger, repo *database.DatabaseManager) *Service {
	wallet := NewWalletService(cfg,logger,repo)
	user := NewUserService(cfg,logger,repo,wallet)	
	return &Service{
		User: user,	
		Wallet: wallet,
	}
}
