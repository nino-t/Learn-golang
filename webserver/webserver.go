package webserver

import (
	"net/http"
	"strconv"

	"github.com/go-learn/pkg/todo"

	controller "github.com/go-learn/webserver/controller"
	"github.com/gorilla/mux"
)

type SERVE_CONFIG struct {
	PORT int
}

var listenAndServe = http.ListenAndServe

func Serve(cfg SERVE_CONFIG, todo todo.ICore) {
	var httpPort = ":" + strconv.Itoa(cfg.PORT)

	envConfig := controller.EnvConfig{}
	router := mux.NewRouter()
	controller.Init(router, envConfig, todo)
	listenAndServe(httpPort, router)
}
