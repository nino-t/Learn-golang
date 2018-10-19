package todo

import (
	myql "github.com/go-sql-driver/mysql"
)

type (
	TodoAttributes struct {
		ID        int    `json:"id" msg:"id"`
		Title     string `json:"title" msg:"title"`
		Completed int8   `json:"completed" msg:"completed"`
		CreatedAt string `json:"created_at" msg:"created_at"`
	}

	TodoDB struct {
		ID        int           `db:"id"`
		Title     string        `db:"title"`
		Completed int8          `db:"completed"`
		CreatedAt string        `db:"created_at"`
		UpdatedAt myql.NullTime `db:"updated_at"`
		DeletedAt myql.NullTime `db:"deleted_at"`
	}
)
