package repository

import (
	"errors"
	"gin-todolist/logger"
	"gin-todolist/model"
	"gin-todolist/todo"
	"gorm.io/gorm"
	"time"
)

type TodoMysqlRepository struct {
	db     *gorm.DB
	logger *logger.CustomLogger
}

func NewMysqlTodoRepository(sql *gorm.DB) todo.Repository {
	return &TodoMysqlRepository{
		db:     sql,
		logger: logger.GetLoggerInstance(),
	}
}

func (r *TodoMysqlRepository) Save(todo *model.Todo) (*model.Todo, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		todo.CreatedAt = time.Now()

		if err := tx.Create(&todo).Error; err != nil {
			r.logger.Error("Error when save todo %+v. caused by: %+v", todo, err)
			return err
		}
		r.logger.Debug("Success insert todo with id : %d", todo.ID)
		return nil
	})

	return todo, err
}

func (r *TodoMysqlRepository) GetOne(id int) (*model.Todo, error) {
	var item *model.Todo

	result := r.db.First(&item, id)

	if result.Error != nil {
		r.logger.Error("Failed to fetch todo item with id %d caused by: %+v", id, result.Error)
		return nil, result.Error
	}

	return item, nil
}

func (r *TodoMysqlRepository) Delete(id int) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		res := tx.Delete(&model.Todo{}, id)
		if res.Error != nil {
			r.logger.Error("Error when delete todo item with id: %d caused by: %+v", id, res.Error)
			return res.Error
		}

		if res.RowsAffected > 0 {
			return nil
		} else {
			return errors.New("no record Found to delete")
		}
	})
}

func (r *TodoMysqlRepository) GetAll() []*model.Todo {
	var todos = make([]*model.Todo, 0)

	result := r.db.Find(&todos)
	if result.Error != nil {
		r.logger.Error("Error when execute Get All Query to database caused by:%+v", result.Error)
	}
	return todos
}
