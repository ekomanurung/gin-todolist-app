package configuration

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	t.Run("success set Log Level", func(t *testing.T) {
		t.Setenv("APP_LOG_LEVEL", "info")
		ConfigureLogLevel()
		assert.Equal(t, zerolog.InfoLevel, zerolog.GlobalLevel())
	})

	t.Run("failed: invalid log level defined", func(t *testing.T) {
		t.Setenv("APP_LOG_LEVEL", "InfoLevel")
		ConfigureLogLevel()
		assert.Equal(t, zerolog.WarnLevel, zerolog.GlobalLevel())
	})
}
