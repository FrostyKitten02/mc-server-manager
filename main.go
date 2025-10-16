package main

import (
	"mc-server-manager/conf"
	"mc-server-manager/httpserver"
	"mc-server-manager/ws"
)

func main() {
	conf.Load()

	go httpserver.StartHttpServer(conf.Conf.ServerPort)
	go ws.StartWsServer(conf.Conf.WsPort)

	for {

	}
}
