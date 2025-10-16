package logs

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/moby/moby/client"
	"io"
	"log/slog"
)

func GetContainerLogs(containerName string, wsConn *websocket.Conn) {
	apiClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		slog.Info("GetContainerLogs error:", err)
		return
	}
	defer apiClient.Close()

	containerLogOptions := client.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
		Since:      "10min",
		Timestamps: true,
		Tail:       "50",
	}
	logs, logsErr := apiClient.ContainerLogs(context.Background(), containerName, containerLogOptions)
	if logsErr != nil {
		slog.Error("Error getting container logs:", logsErr)
		return
	}
	defer logs.Close()

	wsWriter := &WebSocketWriter{Conn: wsConn}
	_, copyErr := io.Copy(wsWriter, logs)
	if copyErr != nil {
		slog.Error("Error copying logs to websocket:", copyErr)
	}

}

type WebSocketWriter struct {
	Conn *websocket.Conn
}

// Write sends model to the WebSocket client
func (w *WebSocketWriter) Write(p []byte) (int, error) {
	err := w.Conn.WriteMessage(websocket.TextMessage, p)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}
