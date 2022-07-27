package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"gin-todolist/logger"
	"gin-todolist/model"
	"gin-todolist/todo"
	"time"
)

type TodoMysqlRepository struct {
	db     *sql.DB
	logger *logger.CustomLogger
}

func NewMysqlTodoRepository(sql *sql.DB) todo.Repository {
	return &TodoMysqlRepository{
		db:     sql,
		logger: logger.GetLoggerInstance(),
	}
}

func (r *TodoMysqlRepository) Save(todo *model.Todo) (*model.Todo, error) {
	todo.CreatedAt = time.Now()

	stmt, err := r.db.Prepare("INSERT INTO todo(title, author, created_at) VALUES(? , ? , ?)")
	if err != nil {
		r.logger.Error("Error when prepare db statement caused by: %+v", err)
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(todo.Title, todo.Author, todo.CreatedAt)
	if err != nil {
		r.logger.Error("Error when save todo %+v. caused by: %+v", todo, err)
		return nil, err
	}

	lastInsertedId, err := res.LastInsertId()
	if err != nil {
		r.logger.Error("Error when get last inserted id caused by: %+v", err)
	}

	todo.ID = int(lastInsertedId)

	r.logger.Debug("Success insert todo with id : %d", lastInsertedId)
	return todo, nil
}

func (r *TodoMysqlRepository) GetOne(id int) (*model.Todo, error) {
	stmt, err := r.db.Prepare("SELECT * FROM todo where id = ?")
	if err != nil {
		r.logger.Error("Error when Prepare Statement Get By id: %d to database caused by: %+v", id, err)
		return nil, err
	}
	defer stmt.Close()

	todo := &model.Todo{}

	err = stmt.QueryRow(id).Scan(&todo.ID, &todo.Title, &todo.Author, &todo.CreatedAt)
	if err != nil {
		r.logger.Error("Error when execute Get By id: %d to database caused by: %+v", id, err)
		return nil, err
	}

	r.logger.Debug("Success Get one with id : %d", id)

	return todo, nil
}

func (r *TodoMysqlRepository) Delete(id int) (bool, error) {
	trx, err := r.db.Begin()
	if err != nil {
		r.logger.Error("Error when create transaction caused by: %+v", err)
		return false, err
	}

	stmt, err := trx.Prepare("DELETE from todo where id = (?)")
	if err != nil {
		r.logger.Error("Error when Prepare Delete Query Statement caused by: %+v", err)
		return false, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		r.logger.Error("Error when Execute Delete Query caused by: %+v", err)
		return false, err
	}

	err = trx.Commit()
	if err != nil {
		r.logger.Error("Error when Commit Delete Query caused by: %+v", err)
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.logger.Error("Error when Fetch Delete Result caused by: %+v", err)
		return false, err
	}

	if rowsAffected > 0 {
		return true, errors.New(fmt.Sprintf("Success delete todo item with id: %d", id))
	}

	return false, errors.New(fmt.Sprintf("Failed to delete todo with id: %d, not found", id))
}

func (r *TodoMysqlRepository) GetAll() []*model.Todo {
	var todos = make([]*model.Todo, 0)

	rows, err := r.db.Query("SELECT * FROM todo")
	if err != nil {
		r.logger.Error("Error when execute Get All Query to database caused by:%+v", err)
		return todos
	}
	defer rows.Close()

	for rows.Next() {
		todo := &model.Todo{}

		err := rows.Scan(&todo.ID, &todo.Title, &todo.Author, &todo.CreatedAt)
		if err != nil {
			r.logger.Error("Error parsing data todo caused by: %+v", err)
		} else {
			todos = append(todos, todo)
		}
	}

	return todos
}
