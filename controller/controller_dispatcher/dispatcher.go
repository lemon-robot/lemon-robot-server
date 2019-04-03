package controller_dispatcher

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-server/controller/http_common"
)

const urlPrefix = "/dispatcher"

func RegApis(router *gin.RouterGroup) {
	router.POST(urlPrefix, register)
}

func register(ctx *gin.Context) {
	http_common.Success(ctx, "pong")
}
