package sysinfo

import (
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/utils/lruio"
	"lemon-robot-server/model"
	"path"
)

var lrConfigObj = &model.LrConfig{}

const configFileName = "lemon.robot.json"

func configFilePath() string {
	return lruio.GetRuntimePath(configFileName)
}

func checkConfigExisted() bool {
	return lruio.PathExists(configFilePath())
}

func LrConfig() *model.LrConfig {
	if checkConfigExisted() {
		err := lruio.JsonToStruct(configFilePath(), &lrConfigObj)
		if err != nil {
			logger.Error("An error occurred while parsing the configuration file_resource, , use default configuration to continue running the system.", err)
			return defaultConfig()
		}
		return lrConfigObj
	}
	logger.Warn("Configuration file_resource not found, use default configuration to continue running the system")
	return defaultConfig()
}

func defaultConfig() *model.LrConfig {
	return &model.LrConfig{
		DbType:        "",
		DbUrl:         "",
		WorkSpacePath: "",
	}
}

func GetWorkspaceSubPath(dirName string) string {
	return path.Join(LrConfig().WorkSpacePath, dirName)
}

func GetHmacKeyBytes() []byte {
	return []byte(LrConfig().SecretHmacKeyword)
}
