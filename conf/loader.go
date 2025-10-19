package conf

import (
	"golang.org/x/oauth2"
	"mc-server-manager/conf/model"
)

var (
	Servers     []model.ServerConfig
	Conf        model.Config
	GoogleOauth oauth2.Config
)

func Load() {
	Servers = readServersConf()
	Conf = readConfig()
	GoogleOauth = readOAuthConfig()
}
