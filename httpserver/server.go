package httpserver

import (
	"github.com/FrostyKitten02/mc-server-manager/httpserver/controller"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

func StartHttpServer(port uint64) {
	router := mux.NewRouter()
	controller.AddAuthMiddleware(router)

	controller.BindAuthEndpoints(router)
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
