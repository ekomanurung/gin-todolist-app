package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)
import "github.com/gin-gonic/gin"

type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}

var todos = []Todo{
	{ID: 1, Title: "Cuci piring", Author: "Eko", CreatedAt: time.Now()},
	{ID: 2, Title: "Masak nasi", Author: "Eko", CreatedAt: time.Now()},
}

func getAllTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func addTodoItem(c *gin.Context) {
	var todo Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	todo.CreatedAt = time.Now()

	todos = append(todos, todo)
	c.JSON(http.StatusCreated, todo)
}

func deleteTodoItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request for param id",
		})
		return
	}

	found := false
	for idx, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:idx], todos[idx+1:]...)
			found = true
			break
		}
	}

	if !found {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("item with id %d not found", id),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("success delete item with id %d", id),
	})
}

func main() {
	r := gin.New()

	r.GET("/todos", getAllTodos)
	r.POST("/todos", addTodoItem)
	r.DELETE("/todos/:id", deleteTodoItem)

	r.Run()
}
