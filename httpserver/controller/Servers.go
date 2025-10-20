package controller

import (
	"encoding/json"
	"github.com/FrostyKitten02/mc-server-manager/conf"
	"github.com/FrostyKitten02/mc-server-manager/conf/model"
	"github.com/gorilla/mux"
	"net/http"
)

func BindServersEndpoints(router *mux.Router) {
	router.HandleFunc("/servers", getAvailableServers).Methods("GET")
}

func getAvailableServers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.MarshalIndent(model.MapServersConfToDto(conf.Servers), "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
