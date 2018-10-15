package webserver

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	controller "github.com/go-learn/webserver/controller"
)

type SERVE_CONFIG struct {
	PORT int
}

var listenAndServe = http.ListenAndServe

func Serve(cfg SERVE_CONFIG) {
	var httpPort = ":" + strconv.Itoa(cfg.PORT)

	envConfig := controller.EnvConfig{}
	router := mux.NewRouter()
	controller.Init(router, envConfig)
	listenAndServe(httpPort, router)
}
