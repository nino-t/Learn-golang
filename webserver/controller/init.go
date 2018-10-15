package controller

import (
	"net/http"

	"github.com/go-learn/webserver/view"
	"github.com/gorilla/mux"
)

type EnvConfig struct{}

type handler struct {
	cfg EnvConfig
}

func Init(r *mux.Router, cfg EnvConfig) {
	h := &handler{
		cfg: cfg,
	}

	r.HandleFunc("/ping", h.handlePing).Methods("GET")
}

func (h *handler) handlePing(w http.ResponseWriter, r *http.Request) {
	view.RenderJSONData(w, "PONG", http.StatusOK)
}
