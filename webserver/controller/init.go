package controller

import (
	"net/http"

	"github.com/go-learn/pkg/todo"

	"github.com/go-learn/webserver/view"
	"github.com/gorilla/mux"
)

type EnvConfig struct{}

type handler struct {
	cfg  EnvConfig
	todo todo.ICore
}

func Init(r *mux.Router, cfg EnvConfig, todo todo.ICore) {
	h := &handler{
		cfg:  cfg,
		todo: todo,
	}

	r.HandleFunc("/ping", h.handlePing).Methods("GET")

	r.HandleFunc("/todos", h.handleTodoList).Methods("GET")
}

func (h *handler) handlePing(w http.ResponseWriter, r *http.Request) {
	view.RenderJSONData(w, "PONG", http.StatusOK)
}

func (h *handler) handleTodoList(w http.ResponseWriter, r *http.Request) {
	res, err := h.getTodoList()
	if err != nil {
		view.RenderJSONError(w, "No video found", http.StatusNotFound)
	} else {
		view.RenderJSONData(w, res, http.StatusOK)
	}

	return
}
