package ws

import (
	"github.com/FrostyKitten02/mc-server-manager/ws/logs"
	"github.com/gorilla/websocket"
	"log/slog"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // For dev only; adjust in prod!
	},
}

func validateJWT(token string) bool {
	return token == "valid_jwt_token"
}

// TODO client need to send server id!!!!
func WsLogsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("Upgrade error:", err)
		return
	}
	defer conn.Close()

	_, msg, err2 := conn.ReadMessage()
	if err2 != nil {
		slog.Error("Error reading JWT token:", err2)
		return
	}
	token := string(msg)
	slog.Debug("Received token:", token)

	if !validateJWT(token) {
		slog.Debug("Invalid token, closing connection")
		err3 := conn.WriteMessage(websocket.TextMessage, []byte("Invalid token"))
		if err3 != nil {
			return
		}
		return
	}

	logs.GetContainerLogs("test", conn)
}

func StartWsServer(port uint64) {
	http.HandleFunc("/logs", WsLogsHandler)

	addr := ":" + strconv.FormatUint(port, 10)
	slog.Info("Starting WebSocket server on " + addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		slog.Error("Error starting websocket server: ", err)
		panic(err)
	}
}
