package conf

import (
	"encoding/json"
	"log/slog"
	"mc-server-manager/conf/model"
	"os"
)

func readServersConf() []model.ServerConfig {
	data, err := os.ReadFile(SERVERS_CONF_LOCATION)
	if err != nil {
		slog.Error("Error reading server conf file", err)
		panic(err)
	}

	var serverConf []model.ServerConfig

	if marshalErr := json.Unmarshal(data, &serverConf); marshalErr != nil {
		slog.Error("Error reading server conf file", marshalErr)
		panic(marshalErr)
	}

	return serverConf
}

func readConfig() model.Config {
	data, err := os.ReadFile(CONFIG_LOCATION)
	if err != nil {
		slog.Error("Error reading config file", err)
		panic(err)
	}

	var conf model.Config
	if marshalErr := json.Unmarshal(data, &conf); marshalErr != nil {
		slog.Error("Error reading config file", err)
		panic(marshalErr)
	}

	return conf
}
