package configuration

import (
	"testing"

	"gin-todolist/model"
	"gin-todolist/todo/handler"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	swaggerFiles "github.com/swaggo/files"
)

func TestApiRegistration(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	t.Run("test todo handler", func(t *testing.T) {
		ApiRegistration(engine, &handler.TodoHandler{})
	})

	t.Run("test swagger handler", func(t *testing.T) {
		ApiRegistration(engine, swaggerFiles.Handler)
	})

	t.Run("panic - test other handler", func(t *testing.T) {
		assert.Panics(t, func() {
			ApiRegistration(engine, &model.Errors{})
		})
	})
}
