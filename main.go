package main

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/utils/lemonrobot"
	"lemon-robot-server/controller"
	"lemon-robot-server/db"
	"lemon-robot-server/sysinfo"
)

func main() {
	startUp()
}

func startUp() {
	logger.Info("Start the lemon-robot startup process")
	lemonrobot.PrintInfo(sysinfo.AppName(), sysinfo.AppVersion())
	db.InitDb()

	engine := gin.Default()
	controller.RegAllApis(engine)

	logger.Info("The system startup process is completed.")
	_ = engine.Run(":33385")
}
