package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ServerDto struct {
	Id             uint64 `json:"id"`
	Name           string `json:"name"`
	McVersion      string `json:"mcVersion"`
	Modpack        string `json:"modpack"`
	ModpackVersion string `json:"modpackVersion"`
	Ip             string `json:"ip"`
	WsConnection   string `json:"wsConnection"`
}

func BindServersEndpoints(router *mux.Router) {
	router.HandleFunc("/servers", getAvailableServers).Methods("GET")
}

func getAvailableServers(w http.ResponseWriter, r *http.Request) {
	servers := []ServerDto{
		{
			Id:             1,
			Name:           "gtnh",
			McVersion:      "1.7.10",
			ModpackVersion: "2.8.0",
			Modpack:        "Greg tech new horizons",
			Ip:             "127.0.0.1",
			WsConnection:   "ws://localhost:8080/servers/1/logs/",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	data, err := json.MarshalIndent(servers, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
