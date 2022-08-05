package configuration

import (
	"testing"

	"gin-todolist/model"
	"gin-todolist/todo/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
)

func TestApiRegistration(t *testing.T) {
	t.Run("test todo handler", func(t *testing.T) {
		ApiRegistration(gin.Default(), &handler.TodoHandler{})
	})

	t.Run("test swagger handler", func(t *testing.T) {
		ApiRegistration(gin.Default(), swaggerFiles.Handler)
	})

	t.Run("panic - test other handler", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("No Panic")
			}
		}()
		ApiRegistration(gin.Default(), &model.Errors{})
	})
}
