package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"gin-todolist/model"
	"gin-todolist/todo"
	"time"
)

type TodoMysqlRepository struct {
	db *sql.DB
}

func NewMysqlTodoRepository(sql *sql.DB) todo.Repository {
	return &TodoMysqlRepository{
		db: sql,
	}
}

func (r *TodoMysqlRepository) Save(todo *model.Todo) (*model.Todo, error) {
	todo.CreatedAt = time.Now()

	stmt, err := r.db.Prepare("INSERT INTO todo(title, author, created_at) VALUES(? , ? , ?)")
	if err != nil {
		fmt.Printf("Error when prepare db statement caused by: %+v\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(todo.Title, todo.Author, todo.CreatedAt)
	if err != nil {
		fmt.Printf("Error when save todo %+v. caused by: %+v\n", todo, err.Error())
		return nil, err
	}

	lastInsertedId, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("Error when get last inserted id caused by: %+v\n", err.Error())
	}

	todo.ID = int(lastInsertedId)
	fmt.Printf("Success insert todo with id : %d\n", lastInsertedId)
	return todo, nil
}

func (r *TodoMysqlRepository) GetOne(id int) (*model.Todo, error) {
	stmt, err := r.db.Prepare("SELECT * FROM todo where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	todo := &model.Todo{}

	err = stmt.QueryRow(id).Scan(&todo.ID, &todo.Title, &todo.Author, &todo.CreatedAt)
	if err != nil {
		fmt.Printf("Error when execute Get By id: %d Query to database caused by: %v\n", id, err.Error())
		return nil, err
	}

	return todo, nil
}

func (r *TodoMysqlRepository) Delete(id int) (bool, error) {
	trx, err := r.db.Begin()
	if err != nil {
		fmt.Sprintf("Error when create transaction caused by: %+v\n", err)
		return false, err
	}

	stmt, err := trx.Prepare("DELETE from todo where id = (?)")
	if err != nil {
		fmt.Sprintf("Error when Prepare Delete Query Statement caused by: %+v\n", err)
		return false, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		fmt.Sprintf("Error when Execute Delete Query caused by: %+v\n", err)
		return false, err
	}

	err = trx.Commit()
	if err != nil {
		fmt.Sprintf("Error when Commit Delete Query caused by: %+v\n", err)
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Sprintf("Error when Fetch Delete Result caused by: %+v\n", err)
		return false, err
	}

	if rowsAffected > 0 {
		return true, errors.New(fmt.Sprintf("Success delete todo item with id: %d", id))
	}

	return false, errors.New(fmt.Sprintf("Failed to delete todo with id: %d, not found", id))
}

func (r *TodoMysqlRepository) GetAll() []*model.Todo {
	rows, err := r.db.Query("SELECT * FROM todo")
	if err != nil {
		fmt.Printf("Error when execute Get All Query to database caused by: %v\n", err.Error())
		return nil
	}
	defer rows.Close()

	var todos = make([]*model.Todo, 0)

	for rows.Next() {
		todo := &model.Todo{}

		err := rows.Scan(&todo.ID, &todo.Title, &todo.Author, &todo.CreatedAt)
		if err != nil {
			fmt.Printf("Error parsing data todo caused by: %+v\n", err)
		}

		todos = append(todos, todo)
	}

	return todos
}
