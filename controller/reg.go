package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/model"
	"lemon-robot-server/controller/controller_file_resource"
	"lemon-robot-server/controller/controller_task"
	"lemon-robot-server/controller/controller_user"
	"lemon-robot-server/controller/http_common"
	"lemon-robot-server/dao/dao_user"
	"lemon-robot-server/entity"
	"lemon-robot-server/sysinfo"
	"net/http"
	"strings"
)

func RegAllApis(engine *gin.Engine) {
	authRouter := engine.Group("/", checkAuthHandler())

	controller_file_resource.RegApis(authRouter)
	controller_user.RegApis(authRouter)
	controller_task.RegApis(authRouter)
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
		Code:    http_common.ErrCode_Common_Unauthorized,
		Data:    nil,
	})
	ctx.Abort()
}

func checkAuthToken(ctx *gin.Context) bool {
	jwtTokenStr := strings.Split(ctx.Request.Header["Authorization"][0], " ")[1]
	jwtToken, err := jwt.Parse(jwtTokenStr, func(token *jwt.Token) (i interface{}, e error) {
		return sysinfo.GetHmacKeyBytes(), nil
	})
	userKey := jwtToken.Claims.(jwt.MapClaims)["iss"]
	user := dao_user.FirstByExample(&entity.User{UserKey: userKey.(string)})
	// user not found or have error
	if user.UserKey == "" || err != nil {
		return false
	}
	if _, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		return true
	} else {
		return false
	}
}
