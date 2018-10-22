package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-learn/pkg/todo"
	"github.com/go-learn/pkg/user"

	"github.com/go-learn/webserver/view"
	"github.com/gorilla/mux"
)

type EnvConfig struct{}

type handler struct {
	cfg  EnvConfig
	todo todo.ICore
	user user.ICore
}

func Init(r *mux.Router, cfg EnvConfig, todo todo.ICore, user user.ICore) {
	h := &handler{
		cfg:  cfg,
		todo: todo,
		user: user,
	}

	r.HandleFunc("/ping", h.handlePing).Methods("GET")

	r.HandleFunc("/todos", h.handleTodoList).Methods("GET")
	r.HandleFunc("/todos", h.handleTodoStore).Methods("POST")
	r.HandleFunc("/todos/{id}", h.handleTodoDetail).Methods("GET")
	r.HandleFunc("/todos/{id}", h.handleTodoUpdate).Methods("PUT")
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

	todo := todo.TodoModel{}
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
	primaryID, _ := strconv.Atoi(queryParams["id"])

	var todo = todo.TodoModel{}
	todo.ID = primaryID

	res, err := h.getTodo(&todo)
	if err != nil {
		view.RenderJSONError(w, "Failed to get todo", http.StatusBadRequest)
	} else {
		view.RenderJSONData(w, res, http.StatusOK)
	}

	return
}

func (h *handler) handleTodoUpdate(w http.ResponseWriter, r *http.Request) {
	queryParams := mux.Vars(r)
	primaryID, _ := strconv.Atoi(queryParams["id"])

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		view.RenderJSONError(w, "Failed to parse http body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	todo := todo.TodoModel{}
	todo.ID = primaryID
	if err := json.Unmarshal(body, &todo); err != nil {
		view.RenderJSONError(w, "Failed to encode entry data", http.StatusBadRequest)
		return
	}

	res, err := h.updateTodo(&todo)
	if err != nil {
		view.RenderJSONError(w, "Failed to update todo", http.StatusBadRequest)
	} else {
		view.RenderJSONData(w, res, http.StatusOK)
	}

	return
}

func (h *handler) handleTodoDelete(w http.ResponseWriter, r *http.Request) {
	queryParams := mux.Vars(r)
	primaryID, _ := strconv.Atoi(queryParams["id"])

	todo := todo.TodoModel{}
	todo.ID = primaryID

	res, err := h.deleteTodo(&todo)
	if err != nil {
		view.RenderJSONError(w, "Failed to delete todo", http.StatusBadRequest)
	} else {
		view.RenderJSONData(w, res, http.StatusOK)
	}

	return
}
