package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/model"
	"lemon-robot-server/controller/controller_auth"
	"lemon-robot-server/controller/controller_file_resource"
	"lemon-robot-server/controller/http_common"
	"net/http"
	"strings"
)

func RegAllApis(engine *gin.Engine) {
	authRouter := engine.Group("/", checkAuthHandler())

	controller_file_resource.RegApis(authRouter)
	controller_auth.RegApis(authRouter)
}

func checkAuthHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqUrlPath := ctx.Request.URL.Path
		if strings.Index(reqUrlPath, http_common.BaseUrlPathPrefixFree) == 0 {
			logger.Debug("The free interface received a network request: " + reqUrlPath)
			ctx.Next()
		} else {
			if checkAuthToken(ctx) {
				ctx.Next()
			} else {
				responseAuthError(ctx)
			}
		}
	}
}

func responseAuthError(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, model.HttpResponse{
		Success: false,
		Code:    http_common.ErrCode_Common_Unauthorized,
		Data:    nil,
	})
	ctx.Abort()
}

func checkAuthToken(ctx *gin.Context) bool {
	token := strings.Split(ctx.Request.Header["Authorization"][0], " ")[1]
	fmt.Println("授权token: " + token)
	return true
}
