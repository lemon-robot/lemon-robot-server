package main

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/utils/lemonrobot"
	"lemon-robot-server/controller"
	"lemon-robot-server/core/gitter"
	"lemon-robot-server/core/websocket"
	"lemon-robot-server/db"
	"lemon-robot-server/service"
	"lemon-robot-server/service_impl"
	"lemon-robot-server/sysinfo"
	"os"
)

func main() {
	startUp()
}

func startUp() {
	var timerService service.TimerService = service_impl.NewTimerServiceImpl()
	var serverNodeService service.ServerNodeService = service_impl.NewServerNodeServiceImpl()

	logger.Info("Start the " + sysinfo.AppName() + " startup process")
	lemonrobot.PrintInfo(sysinfo.AppName(), sysinfo.AppVersion())

	if sysinfo.LrServerConfig().DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	db.InitDb()
	SysSelfRepair()
	timerService.StartTimer()
	serverNodeService.RegisterServerNode()

	engine := gin.Default()
	// start rest api
	controller.RegAllApis(engine)
	// start websocket
	websocket.Serve(engine)

	if _, ok := gitter.SupportedTypes()[sysinfo.LrServerConfig().GitType]; !ok {
		logger.Error("This type of git server ["+sysinfo.LrServerConfig().GitType+"] is not supported", nil)
		os.Exit(1)
	}
	logger.Info("Start trying to establish a connection to git server, git server type: " + sysinfo.LrServerConfig().GitType)
	gitter.ConfigIns(sysinfo.LrServerConfig().GitType, sysinfo.LrServerConfig().GitConfig)

	// git clone test

	//_, cloneErr := git.PlainClone("/Users/lemonit_cn/Downloads/test/golang", false, &git.CloneOptions{
	//	URL:      "https://github.com/src-d/go-git",
	//	Progress: os.Stdout,
	//})
	//if cloneErr != nil {
	//	logger.Error("clone with error", cloneErr)
	//}

	logger.Info("The " + sysinfo.AppName() + " startup process is completed.")
	_ = engine.Run(":33385")
}

func SysSelfRepair() {
	//var sss []*service.SelfRepairService = []
	service_impl.NewUserServiceImpl(sysinfo.LrServerConfig().SecretHmacKeyword).SelfRepair()
}
