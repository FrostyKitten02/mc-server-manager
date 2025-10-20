package conf

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

const serversConfName = "server-conf.json"
const confName = "config.json"
const googleOAuthName = "google-oauth.json"

var ServersConfFileLocation string
var ConfFileLocation string
var GoogleOAuthFileLocation string

var dataDir = "data"

func InitData() {
	setBaseDataDir()
	//TODO: validate that files exist!!!
	ServersConfFileLocation = getDataFileLocation(serversConfName)
	ConfFileLocation = getDataFileLocation(confName)
	GoogleOAuthFileLocation = getDataFileLocation(googleOAuthName)
}

func setBaseDataDir() {
	dataDirEnv := strings.TrimSpace(os.Getenv("MCSM_DATA_DIR"))
	if dataDirEnv == "" {
		slog.Info("MCSM_DATA_DIR variable not set, using default " + dataDir)
		validateDataDir(dataDir)
		return
	}

	validateDataDir(dataDirEnv)
	dataDir = dataDirEnv
	slog.Info(fmt.Sprintf("Using %s for data dir specified in MCSM_DATA_DIR variable", dataDir))
}

// TODO: errors might be a little bit misleading if variable is not set!
func validateDataDir(location string) {
	info, err := os.Stat(location)
	if os.IsNotExist(err) {
		slog.Error(fmt.Sprintf("Directory specified in MCSM_DATA_DIR variable does not exist: %s", location))
		panic(err)
	}

	if !info.IsDir() {
		slog.Error(fmt.Sprintf("Path specified in MCSM_DATA_DIR variable is not a directory: %s", location))
		panic(errors.New("Not a directory: " + location))
	}
}

func getDataFileLocation(file string) string {
	return filepath.Join(dataDir, file)
}
