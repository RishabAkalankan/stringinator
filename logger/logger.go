package logger

import (
	"os"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

// type RequestIdHook struct{}

// func (h RequestIdHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
// 	fmt.Println("hello")
// 	reqId := uuid.New().String()
// 	e.Str("request-id", reqId)
// }

var globalLogger zerolog.Logger
var initalLogger zerolog.Logger

func Initialize() {
	initalLogger = zerolog.New(os.Stderr).With().Str("application", "stringinator").Timestamp().Logger()
	globalLogger = initalLogger
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

func UpdateRequestId() {
	reqId := uuid.New().String()
	globalLogger = initalLogger.With().Str("request-id", reqId).Logger()
}
