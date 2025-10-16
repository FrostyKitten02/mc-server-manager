package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"mc-server-manager/ws/logs"
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
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	_, msg, err2 := conn.ReadMessage()
	if err2 != nil {
		log.Println("Error reading JWT token:", err2)
		return
	}
	token := string(msg)
	log.Println("Received token:", token)

	if !validateJWT(token) {
		log.Println("Invalid token, closing connection")
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
	fmt.Println("Starting WebSocket server on " + addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("Error starting websocket server: ", err)
	}
}
