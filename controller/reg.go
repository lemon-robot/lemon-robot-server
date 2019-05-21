package controller

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/model"
	"lemon-robot-server/controller/controller_dispatcher_machine"
	"lemon-robot-server/controller/controller_dispatcher_tag"
	"lemon-robot-server/controller/controller_file_resource"
	"lemon-robot-server/controller/controller_server_node"
	"lemon-robot-server/controller/controller_task"
	"lemon-robot-server/controller/controller_user"
	"lemon-robot-server/define/http_error_code_define"
	"lemon-robot-server/service"
	"lemon-robot-server/service_impl"
	"net/http"
	"strings"
)

var authService service.AuthService = service_impl.NewAuthServiceImpl()

func RegAllApis(engine *gin.Engine) {
	engine.Use(cors())
	authRouter := engine.Group("/", checkAuthHandler())
	controller_dispatcher_machine.RegApis(authRouter)
	controller_dispatcher_tag.RegApis(authRouter)
	controller_file_resource.RegApis(authRouter)
	controller_server_node.RegApis(authRouter)
	controller_task.RegApis(authRouter)
	controller_user.RegApis(authRouter)
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
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
	return authService.CheckToken(jwtTokenStr)
}
