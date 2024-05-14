package service

import (
	"github.com/JonasBrothers12/BackendChallenge/config"
	"github.com/JonasBrothers12/BackendChallenge/database"
	"github.com/JonasBrothers12/BackendChallenge/model"
	"github.com/rs/zerolog"
)

type WalletService struct {
	config *config.Config
	logger *zerolog.Logger
	repo   *database.DatabaseManager
}

func NewWalletService(cfg *config.Config, logger *zerolog.Logger, repo *database.DatabaseManager) *WalletService {
	return &WalletService{
		config: cfg,
		logger: logger,
		repo:   repo,
	}
}

func (s *WalletService) CreateWalletService(userID int64,alias string,currencyID int64,balance int64) error {
	wallet := &model.WalletUser{
		OwnerID:   		userID,
		Alias: 	   		alias,
		CurrencyID: 	currencyID,
		Balance: 		balance,	
	}

	s.logger.Info().Msgf("creating Wallet for user %d", userID)

	return s.repo.MySQL.Wallet.InsertNewWallet(wallet)
}