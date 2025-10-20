package main

import (
	"github.com/FrostyKitten02/mc-server-manager/conf"
	"github.com/FrostyKitten02/mc-server-manager/httpserver"
	"github.com/FrostyKitten02/mc-server-manager/logs"
	"github.com/FrostyKitten02/mc-server-manager/ws"
)

func main() {
	logs.InitLogger()

	conf.InitData()
	conf.Load()

	go httpserver.StartHttpServer(conf.Conf.ServerPort)
	go ws.StartWsServer(conf.Conf.WsPort)

	for {

	}
}
