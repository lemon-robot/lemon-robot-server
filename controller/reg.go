package controller

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/model"
	"lemon-robot-server/controller/controller_dispatcher"
	"lemon-robot-server/controller/controller_file_resource"
	"lemon-robot-server/controller/controller_task"
	"lemon-robot-server/controller/controller_user"
	"lemon-robot-server/define/http_error_code_define"
	"lemon-robot-server/service/service_auth"
	"net/http"
	"strings"
)

func RegAllApis(engine *gin.Engine) {
	authRouter := engine.Group("/", checkAuthHandler())
	controller_file_resource.RegApis(authRouter)
	controller_user.RegApis(authRouter)
	controller_task.RegApis(authRouter)
	controller_dispatcher.RegApis(authRouter)
}

func checkAuthHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqUrlPath := ctx.Request.URL.Path
		if checkIsLoginUrl(reqUrlPath) {
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

func checkIsLoginUrl(reqUrlPath string) bool {
	return reqUrlPath == "/user/login"
}

func responseAuthError(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, model.HttpResponse{
		Success: false,
		Code:    http_error_code_define.Common_Unauthorized,
		Data:    nil,
	})
	ctx.Abort()
}

func checkAuthToken(ctx *gin.Context) bool {
	jwtTokenStr := strings.Split(ctx.Request.Header["Authorization"][0], " ")[1]
	return service_auth.CheckToken(jwtTokenStr)
}
