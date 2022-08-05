package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseProperties(t *testing.T) {
	t.Run("setup db properties - success", func(t *testing.T) {
		t.Setenv("DB_PORT", "3306")
		t.Setenv("DB_USERNAME", "root")
		t.Setenv("DB_PASSWORD", "root")
		t.Setenv("DB_HOST", "localhost")
		t.Setenv("DB_NAME", "todolist")
		t.Setenv("DB_DRIVER", "mysql")

		properties := NewDatabaseProperties()
		assert.NotNil(t, properties)
		assert.Equal(t, "root", properties.username)
		assert.Equal(t, "root", properties.password)
	})

	t.Run("setup db properties - panic", func(t *testing.T) {
		assert.Panics(t, func() {
			NewDatabaseProperties()
		})
	})
}
