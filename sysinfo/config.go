package sysinfo

import (
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/utils/lruio"
	"lemon-robot-server/model"
	"path"
)

var lrServerConfigObj = &model.LrServerConfig{}

const configFileName = "lemon.robot.json"

func configFilePath() string {
	return lruio.GetRuntimePath(configFileName)
}

func checkConfigExisted() bool {
	return lruio.PathExists(configFilePath())
}

func LrServerConfig() *model.LrServerConfig {
	if checkConfigExisted() {
		err := lruio.JsonToStruct(configFilePath(), &lrServerConfigObj)
		if err != nil {
			logger.Error("An error occurred while parsing the configuration file_resource, , use default configuration to continue running the system.", err)
			return defaultConfig()
		}
		return lrServerConfigObj
	}
	logger.Warn("Configuration file_resource not found, use default configuration to continue running the system")
	return defaultConfig()
}

func defaultConfig() *model.LrServerConfig {
	return &model.LrServerConfig{
		DbType:        "",
		DbUrl:         "",
		WorkSpacePath: "",
	}
}

func GetWorkspaceSubPath(dirName string) string {
	return path.Join(LrServerConfig().WorkSpacePath, dirName)
}

func GetHmacKeyBytes() []byte {
	return []byte(LrServerConfig().SecretHmacKeyword)
}
