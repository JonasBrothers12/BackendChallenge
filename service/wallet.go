package service

import (
	"database/sql"

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

func (s *WalletService) CreateWalletService(wallet *model.WalletUser,tx *sql.Tx) error {
	
	s.logger.Info().Msgf("creating Wallet for user %d", wallet.OwnerID)
	return s.repo.MySQL.Wallet.InsertNewWallet(wallet,tx)
}