package configuration

import (
	"gin-todolist/todo/handler"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/net/webdav"
)

func ApiRegistration(r *gin.Engine, v interface{}) {
	switch h := v.(type) {
	case *handler.TodoHandler:
		registerTodoAPI(r, h)
		break
	case *webdav.Handler:
		registerSwaggerAPI(r, h)
	default:
		panic("Unimplemented Handler, Please Register your handler..")
	}
}

func registerSwaggerAPI(r *gin.Engine, handler *webdav.Handler) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(handler))
}

func registerTodoAPI(r *gin.Engine, todoHandler *handler.TodoHandler) {
	group := r.Group("/v1")
	{
		group.GET("/todos/:id", todoHandler.GetOneTodo)
		group.GET("/todos", todoHandler.GetAllTodos)
		group.POST("/todos", todoHandler.AddTodoItem)
		group.DELETE("/todos/:id", todoHandler.DeleteTodoItem)
	}
}
