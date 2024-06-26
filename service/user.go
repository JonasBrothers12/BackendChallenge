package service

import (
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

	tx,err := s.repo.MySQL.Cli.Begin()

	if err != nil{
		return err
	}

	defer tx.Rollback()

	s.logger.Info().Msgf("creating user %s", user.TaxID)

	valueId,err := s.repo.MySQL.User.InsertNewUser(user,tx)

	if err != nil{
		return err
	}

	currency, _, err := s.repo.MySQL.Currency.SelectByCurrencyCode(tx, model.CurrencyBRL)
	if err != nil {
		return err
	}
	wallet := &model.WalletUser{
		OwnerID: valueId,
		CurrencyID: currency.ID,
		Alias: r.FirstName + " wallet",
	}

	err = s.Wallet.CreateWalletService(wallet,tx)

	if err != nil{
		return err
	}

	if err = tx.Commit(); err != nil {
        return err
    }

	return nil
}

