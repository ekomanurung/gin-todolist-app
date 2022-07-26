package todo

import "gin-todolist/model"

type Repository interface {
	Save(todo *model.Todo) (*model.Todo, error)
	Delete(id int) (bool, error)
	GetOne(id int) (*model.Todo, error)
	GetAll() []*model.Todo
}
