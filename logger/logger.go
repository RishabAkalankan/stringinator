package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var globalLogger zerolog.Logger

func Initialize() {
	logger := zerolog.New(os.Stderr).With().Str("application", "stringinator").Timestamp().Logger()
	globalLogger = logger
	globalLogger.Info().Msg("Logger Initialised")
}

func Infof(message string, args ...interface{}) {
	globalLogger.Info().Msgf(message, args...)
}

func Warnf(message string, args ...interface{}) {
	globalLogger.Warn().Msgf(message, args...)
}

func Errorf(message string, args ...interface{}) {
	globalLogger.Error().Msgf(message, args...)
}

func Fatalf(message string, args ...interface{}) {
	globalLogger.Fatal().Msgf(message, args...)
}
