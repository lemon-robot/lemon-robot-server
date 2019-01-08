package controller

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-server/controller/controller_file_resource"
	"lemon-robot-server/controller/controller_lrc"
)

func RegAllApis(engine *gin.Engine) {
	controller_file_resource.RegApis(engine)
	controller_lrc.RegApis(engine)
}
