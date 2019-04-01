package controller_dispatcher

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-server/controller/http_common"
)

const urlPrefix = "/dispatcher"
const urlPing = "/ping"

func RegApis(router *gin.RouterGroup) {
	router.POST(urlPrefix+urlPing, ping)
}

func ping(ctx *gin.Context) {
	http_common.Success(ctx, "pong")
}
