package main

import (
	"database/sql"
	"fmt"
	"gin-todolist/logger"
	"gin-todolist/todo/handler"
	"gin-todolist/todo/repository"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	r := gin.New()

	logger.SetLogLevel(2)

	//initialize mysql configuration
	//TODO remove parseTime attribute, since we should define the date time in db in a right way
	db, sqlErr := sql.Open("mysql", "root:root@tcp(mysql:3306)/todolist?parseTime=true")
	if sqlErr != nil {
		panic(fmt.Sprintf("Panic when initialize mysql connection caused by: %+v\n", sqlErr.Error()))
	}

	defer db.Close()

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 20)

	todoRepository := repository.NewMysqlTodoRepository(db)
	handler.NewTodoHandler(r, todoRepository)

	err := r.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Panic when starting the web server caused by: %v\n", err.Error()))
	}
}
