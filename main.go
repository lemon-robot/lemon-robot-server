package main

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-golang-commons/utils/lemonrobot"
	"lemon-robot-server/controller"
	"lemon-robot-server/db"
	"lemon-robot-server/sysinfo"
)

func main() {
	startUp()
}

func startUp() {
	lemonrobot.PrintInfo(sysinfo.AppName(), sysinfo.AppVersion())
	db.InitDb()

	engine := gin.Default()
	controller.RegAllApis(engine)
	_ = engine.Run()
}
