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
	r.HandleFunc("/todos/{id}", h.handleTodoDetail).Methods("GET")
	r.HandleFunc("/todos/{id}", h.handleTodoDelete).Methods("DELETE")
}

func (h *handler) handlePing(w http.ResponseWriter, r *http.Request) {
	view.RenderJSONData(w, "PONG", http.StatusOK)
}

func (h *handler) handleTodoList(w http.ResponseWriter, r *http.Request) {
	res, err := h.getTodoList()
	if err != nil {
		view.RenderJSONError(w, "No todo found", http.StatusNotFound)
	} else {
		view.RenderJSONData(w, res, http.StatusOK)
	}

	return
}

func (h *handler) handleTodoStore(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		view.RenderJSONError(w, "Failed to parse http body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	todo := todo.TodoData{}
	if err := json.Unmarshal(body, &todo); err != nil {
		view.RenderJSONError(w, "Failed to encode entry data", http.StatusBadRequest)
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

func (h *handler) handleTodoDetail(w http.ResponseWriter, r *http.Request) {
	queryParams := mux.Vars(r)
	primaryId := queryParams["id"]

	res, err := h.getTodo(primaryId)
	if err != nil {
		view.RenderJSONError(w, "Failed to get todo", http.StatusBadRequest)
	} else {
		view.RenderJSONData(w, res, http.StatusOK)
	}

	return
}

func (h *handler) handleTodoDelete(w http.ResponseWriter, r *http.Request) {
	queryParams := mux.Vars(r)
	primaryId := queryParams["id"]

	res, err := h.deleteTodo(primaryId)
	if err != nil {
		view.RenderJSONError(w, "Failed to delete todo", http.StatusBadRequest)
	} else {
		view.RenderJSONData(w, res, http.StatusOK)
	}

	return
}
