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
}

func NewUserService(cfg *config.Config, logger *zerolog.Logger, repo *database.DatabaseManager) *UserService {
	return &UserService{
		config: cfg,
		logger: logger,
		repo:   repo,
	}
}

func (s *UserService) CreateUserService(r *requisition.UserAccountRequest) error {
	user := &model.UserAcount{
		FirstName:      r.FirstName,
		LastName:       r.LastName,
		TaxID: 			r.TaxID,
	}

	s.logger.Info().Msgf("creating user %s", user.TaxID)

	return s.repo.MySQL.User.InsertNewUser(user)
}