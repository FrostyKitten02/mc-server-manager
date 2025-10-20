package conf

import (
	"encoding/json"
	"github.com/FrostyKitten02/mc-server-manager/conf/model"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log/slog"
	"os"
)

func readServersConf() []model.ServerConfig {
	data, err := os.ReadFile(ServersConfFileLocation)
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
	data, err := os.ReadFile(ConfFileLocation)
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

func readOAuthConfig() oauth2.Config {
	data, err := os.ReadFile(GoogleOAuthFileLocation)
	if err != nil {
		slog.Error("Error reading config file", err)
		panic(err)
	}

	var conf model.GoogleOauth
	if marshalErr := json.Unmarshal(data, &conf); marshalErr != nil {
		slog.Error("Error reading google oauth config file", err)
		panic(marshalErr)
	}

	endpoint := google.Endpoint
	endpoint.AuthStyle = oauth2.AuthStyleAutoDetect
	return oauth2.Config{
		ClientID:     conf.Web.ClientId,
		ClientSecret: conf.Web.ClientSecret,
		RedirectURL:  "http://localhost:4000/callback", //TODO!!
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: endpoint,
	}
}
