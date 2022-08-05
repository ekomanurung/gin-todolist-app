package configuration

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func ConfigureLogLevel() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logLevel := os.Getenv("APP_LOG_LEVEL")

	l, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		log.Error().Msgf("Switch log level to Warn, Because parse Log Level Failed, caused by %+v", err)
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	} else {
		zerolog.SetGlobalLevel(l)
	}
}
