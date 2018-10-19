package controller

import (
	"log"

	"github.com/go-learn/pkg/todo"
	"github.com/vmihailenco/msgpack"
)

func (h *handler) getTodoList() ([]todo.TodoAttributes, error) {
	collection := make([]todo.TodoAttributes, 0)

	result, err := h.todo.GetTodoListFromDB()
	if err != nil {
		log.Printf("[Get Todos] Failed To Get Todos, Error: %v", err)
		return collection, err
	}

	for _, data := range result {
		var TodoAttributes todo.TodoAttributes
		todoAttributesByte, err := msgpack.Marshal(data)

		err = msgpack.Unmarshal(todoAttributesByte, &TodoAttributes)
		if err != nil {
			log.Printf("[Get Todos] Failed to Unmarshal Todo Attributes %d, Error: %v", data.ID, err)
			return collection, err
		}

		collection = append(collection, TodoAttributes)
	}

	return collection, nil
}
