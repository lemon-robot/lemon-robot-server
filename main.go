package main

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/utils/lemonrobot"
	"lemon-robot-server/controller"
	"lemon-robot-server/core/git"
	"lemon-robot-server/db"
	"lemon-robot-server/service/service_user"
	"lemon-robot-server/sysinfo"
	"os"
)

func main() {
	startUp()
}

func startUp() {
	logger.Info("Start the lemon-robot startup process")
	lemonrobot.PrintInfo(sysinfo.AppName(), sysinfo.AppVersion())

	if sysinfo.LrConfig().DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	db.InitDb()
	SysSelfRepair()

	engine := gin.Default()
	controller.RegAllApis(engine)

	if _, ok := git.SupportedTypes()[sysinfo.LrConfig().GitType]; !ok {
		logger.Error("This type of git server ["+sysinfo.LrConfig().GitType+"] is not supported", nil)
		os.Exit(1)
	}
	logger.Info("Start trying to establish a connection to git server, git server type: " + sysinfo.LrConfig().GitType)
	var gitIns git.Standard
	gitIns = git.SupportedTypes()[sysinfo.LrConfig().GitType]
	gitIns.Init(sysinfo.LrConfig().GitConfig)

	logger.Info("The " + sysinfo.AppName() + " startup process is completed.")
	_ = engine.Run(":33385")
}

func SysSelfRepair() {
	service_user.SelfRepair()
}
