package main

import (
	"github.com/gorilla/mux"
	"mc-server-manager/controller"
	"mc-server-manager/ws"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	controller.BindServersEndpoints(router)
	err := http.ListenAndServe(":4000", router)
	if err != nil {
		return
	}

	ws.StartWsServer()

	for {

	}
}
