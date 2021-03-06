package model

import "time"

type Todo struct {
	ID        int       `db:"id" json:"id"`
	Title     string    `db:"title" json:"title" binding:"required"`
	Author    string    `db:"author" json:"author" binding:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
