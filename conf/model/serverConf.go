package model

import "strconv"

type ServerDto struct {
	Id             uint64 `json:"id"`
	Name           string `json:"name"`
	McVersion      string `json:"mcVersion"`
	Modpack        string `json:"modpack"`
	ModpackVersion string `json:"modpackVersion"`
	LogsUrl        string `json:"logsUrl"`
}

type ServerConfig struct {
	Id             uint64 `json:"id"`
	DockerName     string `json:"dockerName"`
	Name           string `json:"name"`
	McVersion      string `json:"mcVersion"`
	Modpack        string `json:"modpack"`
	ModpackVersion string `json:"modpackVersion"`
	Ip             string `json:"ip"`
}

func MapServerConfToDto(conf ServerConfig) ServerDto {
	return ServerDto{
		Id:             conf.Id,
		Name:           conf.Name,
		McVersion:      conf.McVersion,
		Modpack:        conf.Modpack,
		ModpackVersion: conf.ModpackVersion,
		LogsUrl:        strconv.FormatUint(conf.Id, 10) + "/logs", //TODO: get ws server url and add it, this is this servers ip!!
	}
}

func MapServersConfToDto(conf []ServerConfig) []ServerDto {
	result := make([]ServerDto, len(conf))
	for i, server := range conf {
		result[i] = MapServerConfToDto(server)
	}
	return result
}
