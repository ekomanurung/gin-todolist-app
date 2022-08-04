package main

import (
	"database/sql"
	"fmt"
	"time"

	"gin-todolist/docs"
	"gin-todolist/model"
	"gin-todolist/todo/handler"
	"gin-todolist/todo/repository"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeMysql() *gorm.DB {
	sourceName := "root:root@tcp(localhost:3306)/todolist?parseTime=true"

	conn, err := sql.Open("mysql", sourceName)
	if err != nil {
		panic(fmt.Sprintf("Panic when initialize mysql driver connection caused by: %+v\n", err))
	}

	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)
	conn.SetConnMaxLifetime(time.Minute * 20)

	db, err := gorm.Open(mysql.New(mysql.Config{Conn: conn}))
	if err != nil {
		panic(fmt.Sprintf("Panic when initialize Gorm db caused by: %+v\n", err))
	}

	//auto migrate table todos
	db.AutoMigrate(&model.Todo{})

	return db
}

func initializeLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.WarnLevel)
}

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	r := gin.Default()

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Todolist API"
	docs.SwaggerInfo.Description = "Go todolist service using swagger"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	//Initialize mysql configuration
	db := initializeMysql()

	//Initialize Logger
	initializeLogger()

	todoRepository := repository.NewMysqlTodoRepository(db)
	handler.NewTodoHandler(r, todoRepository)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Panic when starting the web server caused by: %v\n", err.Error()))
	}
}
