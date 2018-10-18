package controller

import (
	"log"
)

func (h *handler) getTodoList() (interface{}, error) {
	res, err := h.todo.GetTodoListFromDB()
	if err != nil {
		log.Printf("[Get Video] Failed To Get Video Token, Error: %v", err)
		return res, err
	}

	return res, nil
}
