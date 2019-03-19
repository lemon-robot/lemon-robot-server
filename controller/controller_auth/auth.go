package controller_auth

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-server/controller/http_common"
)

const urlPrefix = "/auth"
const urlLrc = "/lr_user"

func RegApis(router *gin.RouterGroup) {
	router.POST(http_common.BaseUrlPathPrefixFree+urlPrefix+urlLrc, authLrUser)

	router.PUT(urlPrefix+urlLrc, createLrUser)
}
