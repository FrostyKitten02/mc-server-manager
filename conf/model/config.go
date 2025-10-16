package model

type Config struct {
	ServerIp   string `json:"serverIp"`
	ServerPort uint64 `json:"serverPort"`
	WsPort     uint64 `json:"wsPort"`
}
