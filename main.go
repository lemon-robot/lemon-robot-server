package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/src-d/go-git.v4"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/utils/lemonrobot"
	"lemon-robot-server/controller"
	"lemon-robot-server/core/gitter"
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

	if _, ok := gitter.SupportedTypes()[sysinfo.LrConfig().GitType]; !ok {
		logger.Error("This type of gitter server ["+sysinfo.LrConfig().GitType+"] is not supported", nil)
		os.Exit(1)
	}
	logger.Info("Start trying to establish a connection to gitter server, gitter server type: " + sysinfo.LrConfig().GitType)
	gitter.ConfigIns(sysinfo.LrConfig().GitType, sysinfo.LrConfig().GitConfig)

	_, cloneErr := git.PlainClone("/Users/lemonit_cn/Downloads/test/golang", false, &git.CloneOptions{
		URL:      "https://github.com/src-d/go-git",
		Progress: os.Stdout,
	})
	if cloneErr != nil {
		logger.Error("clone with error", cloneErr)
	}

	logger.Info("The " + sysinfo.AppName() + " startup process is completed.")
	_ = engine.Run(":33385")
}

func SysSelfRepair() {
	service_user.SelfRepair()
}
