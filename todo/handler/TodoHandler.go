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

	group := r.Group("/v1")
	{
		group.GET("/todos/:id", handler.GetOneTodo)
		group.GET("/todos", handler.GetAllTodos)
		group.POST("/todos", handler.AddTodoItem)
		group.DELETE("/todos/:id", handler.DeleteTodoItem)
	}
}

func (t *TodoHandler) AddTodoItem(c *gin.Context) {
	var item *model.Todo

	if err := c.ShouldBindJSON(&item); err != nil {
		model.ValidateStruct(c, err)
		return
	}

	result, err := t.Repository.Save(item)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

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

	err = t.Repository.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, true)
}

func (t *TodoHandler) GetOneTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	item, err := t.Repository.GetOne(id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, item)
}
