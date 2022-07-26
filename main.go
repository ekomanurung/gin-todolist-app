package main

import (
	"database/sql"
	"fmt"
	"gin-todolist/todo/handler"
	"gin-todolist/todo/repository"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	r := gin.New()

	//initialize mysql configuration
	//TODO remove parseTime attribute, since we should define the date time in db in a right way
	db, sqlErr := sql.Open("mysql", "root:root@/todolist?parseTime=true")
	if sqlErr != nil {
		panic(fmt.Sprintf("Panic when initialize mysql connection caused by: %+v\n", sqlErr.Error()))
	}

	//defer db.Close()

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 20)

	todoRepository := repository.NewMysqlTodoRepository(db)
	handler.NewTodoHandler(r, todoRepository)

	err := r.Run()
	if err != nil {
		panic(fmt.Sprintf("Panic when starting the web server caused by: %v\n", err.Error()))
	}
}
