package service

import (
	constants "github.com/JonasBrothers12/BackendChallenge/Constants"
	"github.com/JonasBrothers12/BackendChallenge/config"
	"github.com/JonasBrothers12/BackendChallenge/database"
	"github.com/JonasBrothers12/BackendChallenge/model"
	"github.com/JonasBrothers12/BackendChallenge/presenter/requisition"
	"github.com/rs/zerolog"
)

type UserService struct {
	config *config.Config
	logger *zerolog.Logger
	repo   *database.DatabaseManager
	Wallet *WalletService
}

func NewUserService(cfg *config.Config, logger *zerolog.Logger, repo *database.DatabaseManager,wallet *WalletService) *UserService {
	return &UserService{
		config: cfg,
		logger: logger,
		repo:   repo,
		Wallet: wallet,
	}
}

func (s *UserService) CreateUserService(r *requisition.UserAccountRequest) error {
	user := &model.UserAcount{
		FirstName:      r.FirstName,
		LastName:       r.LastName,
		TaxID: 			r.TaxID,
	}
	s.logger.Info().Msgf("creating user %s", user.TaxID)
	valueId,err := s.repo.MySQL.User.InsertNewUser(user)
	if err != nil{
		return nil
	}
	Alias := r.FirstName + " wallet"
	err = s.Wallet.CreateWalletService(valueId,Alias,constants.BRL,0)
	if err != nil{
		return nil
	}
	return nil
}