package conf

import "mc-server-manager/conf/model"

var (
	Servers []model.ServerConfig
	Conf    model.Config
)

func Load() {
	Servers = readServersConf()
	Conf = readConfig()
}
