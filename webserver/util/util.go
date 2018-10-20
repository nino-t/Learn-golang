package util

import (
	"github.com/go-learn/pkg/todo"
)

func TodoDataToTodoDB(todoData todo.TodoData) todo.TodoDB {
	todo := todo.TodoDB{
		Title: todoData.Title,
	}

	return todo
}
