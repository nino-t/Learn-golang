package webserver

import (
	"net/http"
	"strconv"

	"github.com/go-learn/pkg/todo"
	"github.com/go-learn/pkg/user"

	controller "github.com/go-learn/webserver/controller"
	"github.com/gorilla/mux"
)

type SERVE_CONFIG struct {
	PORT int
}

var listenAndServe = http.ListenAndServe

func Serve(cfg SERVE_CONFIG, todo todo.ICore, user user.ICore) {
	var httpPort = ":" + strconv.Itoa(cfg.PORT)

	envConfig := controller.EnvConfig{}
	router := mux.NewRouter()
	controller.Init(router, envConfig, todo, user)
	listenAndServe(httpPort, router)
}
