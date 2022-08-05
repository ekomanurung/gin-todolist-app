package handler

import (
	"net/http"
	"strconv"

	"gin-todolist/model"
	"gin-todolist/todo"
	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	Repository todo.Repository
}

func NewTodoHandler(repository todo.Repository) *TodoHandler {
	return &TodoHandler{
		Repository: repository,
	}
}

// AddTodo godoc
// @Summary      Add Todo Item
// @Description  create new todo and save to db
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        todo   body    model.Todo true "add todo item"
// @Success      201  {object}  model.Todo
// @Failure      422  {object}  model.Response
// @Failure      400  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /todos [post]
func (t *TodoHandler) AddTodoItem(c *gin.Context) {
	var item *model.Todo

	if err := c.ShouldBindJSON(&item); err != nil {
		model.ValidateStruct(c, err)
		return
	}

	result, err := t.Repository.Save(item)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.Response{
			Code:   http.StatusUnprocessableEntity,
			Errors: []model.Errors{{Message: err.Error()}},
		})
		return
	}

	c.JSON(http.StatusCreated, result)
}

// ListTodos godoc
// @Summary      List Todos
// @Description  get all todos
// @Tags         todos
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.Todo
// @Failure      400  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /todos [get]
func (t *TodoHandler) GetAllTodos(c *gin.Context) {
	c.JSON(http.StatusOK, t.Repository.GetAll())
}

// DeleteTodo godoc
// @Summary      Delete Todo item
// @Description  Delete todo by todo id
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id   path      int        true "Todo Id"
// @Success      200  {array}   model.Todo
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /todos/{id} [delete]
func (t *TodoHandler) DeleteTodoItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:   http.StatusBadRequest,
			Errors: []model.Errors{{Message: err.Error()}},
		})
		return
	}

	err = t.Repository.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:   http.StatusBadRequest,
			Errors: []model.Errors{{Message: err.Error()}},
		})
		return
	}

	c.JSON(http.StatusOK, true)
}

// GetOne godoc
// @Summary      Get one todo
// @Description  get one todo by id
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param         id   path      int        true "Todo Id"
// @Success      200  {array}   model.Todo
// @Failure      404  {object}  model.Response
// @Failure      400  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /todos/{id} [get]
func (t *TodoHandler) GetOneTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:   http.StatusBadRequest,
			Errors: []model.Errors{{Message: err.Error()}},
		})
		return
	}

	item, err := t.Repository.GetOne(id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.Response{
			Code:   http.StatusUnprocessableEntity,
			Errors: []model.Errors{{Message: err.Error()}},
		})
		return
	}
	c.JSON(http.StatusOK, item)
}
