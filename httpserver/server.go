package httpserver

import (
	"github.com/gorilla/mux"
	"log/slog"
	"mc-server-manager/httpserver/controller"
	"net/http"
	"strconv"
	"time"
)

func StartHttpServer(port uint64) {
	router := mux.NewRouter()
	controller.BindServersEndpoints(router)
	addr := "127.0.0.1:" + strconv.FormatUint(port, 10)
	slog.Info("Starting http server on " + addr)
	server := http.Server{
		Handler:      router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		slog.Error("Failed to start http server on "+addr, err)
		panic(err)
	}
}
