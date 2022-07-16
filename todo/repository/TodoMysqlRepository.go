package repository

import (
	"errors"
	"fmt"
	"gin-todolist/model"
	"gin-todolist/todo"
	"time"
)

type TodoMysqlRepository struct {
}

func NewMysqlTodoRepository() todo.Repository {
	return &TodoMysqlRepository{}
}

var todos = []model.Todo{
	{ID: 1, Title: "Cuci piring", Author: "Eko", CreatedAt: time.Now()},
	{ID: 2, Title: "Masak nasi", Author: "Eko", CreatedAt: time.Now()},
}

func (r *TodoMysqlRepository) Save(todo model.Todo) (model.Todo, error) {
	todo.CreatedAt = time.Now()

	todos = append(todos, todo)

	return todo, nil
}

func (r *TodoMysqlRepository) GetOne(id int) (*model.Todo, error) {
	idx := findItemFromDB(id)

	if idx == -1 {
		return nil, errors.New(fmt.Sprintf("item with id %d not found", id))
	}

	return &todos[idx], nil
}

func (r *TodoMysqlRepository) Delete(id int) (bool, error) {
	idx := findItemFromDB(id)

	if idx == -1 {
		return false, errors.New(fmt.Sprintf("item with id %d not found", id))
	}

	todos = append(todos[:idx], todos[idx+1:]...)
	return true, errors.New(fmt.Sprintf("success delete item with id %d", id))
}

func (r *TodoMysqlRepository) GetAll() []model.Todo {
	return todos
}

func findItemFromDB(id int) int {
	for idx, t := range todos {
		if t.ID == id {
			return idx
		}
	}
	return -1
}
