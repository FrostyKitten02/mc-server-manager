package main

import (
	"mc-server-manager/controller"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	RconPort      uint32 `json:"rconPort"`
	Ip            string `json:"ip"`
	ContainerName string `json:"containerName"`
}

func main() {
	router := mux.NewRouter()
	controller.BindServersEndpoints(router)
	err := http.ListenAndServe(":4000", router)
	if err != nil {
		return
	}

	for {

	}
}
