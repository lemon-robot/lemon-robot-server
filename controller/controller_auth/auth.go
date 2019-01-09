package controller_auth

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-server/controller/http_common"
)

const urlPrefix = "/auth"

func RegApis(router *gin.RouterGroup) {
	router.POST(http_common.BaseUrlPathPrefixFree+urlPrefix+"/lrc", authLrc)
}
