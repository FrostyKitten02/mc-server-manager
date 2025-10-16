package conf

import (
	"encoding/json"
	"fmt"
	"mc-server-manager/conf/model"
	"os"
)

func readServersConf() []model.ServerConfig {
	data, err := os.ReadFile(SERVERS_CONF_LOCATION)
	if err != nil {
		fmt.Println("Error reading server conf file")
		panic(err)
	}

	var serverConf []model.ServerConfig

	if marshalErr := json.Unmarshal(data, &serverConf); marshalErr != nil {
		fmt.Println("Error reading server conf file")
		panic(marshalErr)
	}

	return serverConf
}

func readConfig() model.Config {
	data, err := os.ReadFile(CONFIG_LOCATION)
	if err != nil {
		fmt.Println("Error reading config file")
		panic(err)
	}

	var conf model.Config
	if marshalErr := json.Unmarshal(data, &conf); marshalErr != nil {
		fmt.Println("Error reading config file")
		panic(marshalErr)
	}

	return conf
}
