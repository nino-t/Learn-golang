package controller

import (
	"log"

	"github.com/go-learn/pkg/todo"
	"github.com/vmihailenco/msgpack"
)

func (h *handler) getTodoList() ([]todo.TodoModel, error) {
	collection := make([]todo.TodoModel, 0)

	result, err := h.todo.GetTodoListFromDB()
	if err != nil {
		log.Printf("[Get Todos] Failed To Get Todos, Error: %v", err)
		return collection, err
	}

	for _, data := range result {
		var TodoAttributes todo.TodoModel
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

func (h handler) createTodo(todoData *todo.TodoModel) (interface{}, error) {
	result, err := h.todo.CreateTodoFromDB(todoData)
	if err != nil {
		log.Printf("[Create Todo] Failed to create todo, Error: %v", err)
		return nil, err
	}

	return result, nil
}

func (h *handler) getTodo(todoData *todo.TodoModel) (interface{}, error) {
	result, err := h.todo.GetTodoDetailFromDB(todoData)
	if err != nil {
		log.Printf("[Get Todo] Failed To Get Todo, Error: %v", err)
		return nil, err
	}

	return result, nil
}

func (h *handler) updateTodo(todoData *todo.TodoModel) (interface{}, error) {
	result, err := h.todo.UpdateTodoFromDB(todoData)
	if err != nil {
		log.Printf("[Update Todo] Failed to update todo, Error: %v", err)
		return nil, err
	}

	return result, nil
}

func (h *handler) deleteTodo(todoData *todo.TodoModel) (interface{}, error) {
	_, err := h.todo.DeleteTodoFromDB(todoData)
	if err != nil {
		log.Printf("[Delete Todo] Failed to delete todo, Error: %v", err)
		return nil, err
	}

	return nil, nil
}
