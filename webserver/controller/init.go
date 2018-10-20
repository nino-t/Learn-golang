package controller

import (
	"encoding/json"
	"io/ioutil"
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
	r.HandleFunc("/todos", h.handleTodoStore).Methods("POST")
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

func (h *handler) handleTodoStore(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		view.RenderJSONError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	todo := todo.TodoData{}
	if err := json.Unmarshal(body, &todo); err != nil {
		view.RenderJSONError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	res, err := h.createTodo(&todo)
	if err != nil {
		view.RenderJSONError(w, "Failed to insert todo", http.StatusBadRequest)
	} else {
		view.RenderJSONData(w, res, http.StatusOK)
	}

	return
}
