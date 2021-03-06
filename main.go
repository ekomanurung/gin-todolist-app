package main

import (
	"database/sql"
	"fmt"
	"gin-todolist/model"
	"gin-todolist/todo/handler"
	"gin-todolist/todo/repository"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func initializeMysql() *gorm.DB {
	sourceName := "root:root@tcp(mysql:3306)/todolist?parseTime=true"

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

func main() {
	r := gin.Default()

	//Initialize mysql configuration
	db := initializeMysql()

	//Initialize Logger
	initializeLogger()

	todoRepository := repository.NewMysqlTodoRepository(db)
	handler.NewTodoHandler(r, todoRepository)

	err := r.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Panic when starting the web server caused by: %v\n", err.Error()))
	}
}
