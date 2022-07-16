package main

import (
	"gin-todolist/todo/handler"
	"gin-todolist/todo/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	todoRepository := repository.NewMysqlTodoRepository()
	handler.NewTodoHandler(r, todoRepository)

	r.Run()
}
