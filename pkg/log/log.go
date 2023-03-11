package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

var App zerolog.Logger
var Http zerolog.Logger
var Mongo zerolog.Logger
var Redis zerolog.Logger

const (
	ApplicationLogger = "application"
	HTTPLogger        = "http"
	MongoLogger       = "mongo"
	RedisLogger       = "redis"
)

func createLogger(loggerType string) zerolog.Logger {
	return log.With().Str("type", loggerType).Logger()
}

func Initialize(appName string) {
	// set default logger configurations
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	// set the default logger level
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// add service name and filename to the default logger
	log.Logger = log.With().Caller().Str("name", appName).Logger()

	App = createLogger(ApplicationLogger)
	Http = createLogger(HTTPLogger)
	Mongo = createLogger(MongoLogger)
	Redis = createLogger(RedisLogger)
}
