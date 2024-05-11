package logger

import (
	"os"
	"github.com/JonasBrothers12/BackendChallenge/config"
	"github.com/rs/zerolog"
)

func Initiallogger(cfg *config.InternalConfig) *zerolog.Logger {
	logger1 := zerolog.New(os.Stderr).With().Str("Service Name:", cfg.ServiceName).Timestamp().Logger()
	return &logger1
}