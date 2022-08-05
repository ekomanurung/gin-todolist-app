package configuration

import (
	"gin-todolist/model"
	"gin-todolist/todo/handler"
	"gin-todolist/todo/repository"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
)

func ConfigureDependencies(r *gin.Engine) {
	// Database Initialization
	databaseProperties := NewDatabaseProperties()
	dbEngine := NewDatabase()
	dbEngine.InitializeConnection(databaseProperties)
	dbEngine.Migrate(&model.Todo{})

	todoRepository := repository.NewMysqlTodoRepository(dbEngine.GetDBEngine().db)

	//Handler Initialization
	todoHandler := handler.NewTodoHandler(todoRepository)

	//API Registration
	ApiRegistration(r, todoHandler)
	ApiRegistration(r, swaggerFiles.Handler)
}
