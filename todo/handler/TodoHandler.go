package handler

import (
	"gin-todolist/model"
	"gin-todolist/todo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TodoHandler struct {
	Repository todo.Repository
}

func NewTodoHandler(r *gin.Engine, repository todo.Repository) {
	handler := &TodoHandler{
		Repository: repository,
	}

	r.GET("/todos", handler.GetAllTodos)
	r.POST("/todos", handler.AddTodoItem)
	r.DELETE("/todos/:id", handler.DeleteTodoItem)
}

func (t *TodoHandler) AddTodoItem(c *gin.Context) {
	var todo model.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, _ := t.Repository.Save(todo)

	c.JSON(http.StatusCreated, result)
}

func (t *TodoHandler) GetAllTodos(c *gin.Context) {
	c.JSON(http.StatusOK, t.Repository.GetAll())
}

func (t *TodoHandler) DeleteTodoItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request for param id",
		})
		return
	}

	success, err := t.Repository.Delete(id)

	if !success {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": err.Error(),
	})
}
