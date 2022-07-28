package todo

import "gin-todolist/model"

type Repository interface {
	Save(todo *model.Todo) (*model.Todo, error)
	Delete(id int) error
	GetOne(id int) (*model.Todo, error)
	GetAll() []*model.Todo
}
