package conf

import (
	"github.com/FrostyKitten02/mc-server-manager/conf/model"
	"golang.org/x/oauth2"
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
