package controller_user

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-server/controller/http_common"
	"lemon-robot-server/define/http_error_code_define"
	"lemon-robot-server/dto"
	"lemon-robot-server/service"
	"lemon-robot-server/service_impl"
	"lemon-robot-server/sysinfo"
)

var userService service.UserService = service_impl.NewUserServiceImpl(sysinfo.LrServerConfig().SecretHmacKeyword)
var authService service.AuthService = service_impl.NewAuthServiceImpl()

const urlPrefix = "/user"
const urlLogin = "/login"

func RegApis(router *gin.RouterGroup) {
	router.POST(urlPrefix+urlLogin, login)

	router.POST(urlPrefix, create)
}

func login(ctx *gin.Context) {
	reqAuthUser := dto.LrUserAuthReq{}
	http_common.DealError(ctx, ctx.BindJSON(&reqAuthUser), "", func(ctx *gin.Context) {
		result, userEntity := userService.CheckPassword(reqAuthUser.Number, reqAuthUser.Password)
		if result {
			http_common.Success(ctx, authService.GenerateJwtTokenStr(userEntity.UserKey))
		} else {
			http_common.Failed(ctx, http_error_code_define.User_LoginIdentityInfoVerifyFailed)
		}
	})
}

func create(ctx *gin.Context) {
	createUserReq := dto.LrUserCreateReq{}
	http_common.DealError(ctx, ctx.Bind(&createUserReq), "", func(ctx *gin.Context) {
		count := userService.CountByNumber(createUserReq.Number)
		if count > 0 {
			http_common.Failed(ctx, http_error_code_define.User_CreateFailedNumberAlreadyExists)
			return
		}
		err, _ := userService.Create(createUserReq.Number, createUserReq.Password)
		if err == nil {
			http_common.Success(ctx, true)
		} else {
			http_common.Failed(ctx, err.Error())
		}
	})
}
