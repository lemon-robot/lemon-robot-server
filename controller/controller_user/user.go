package controller_user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lemon-robot-server/controller/http_common"
	"lemon-robot-server/dao/dao_user"
	"lemon-robot-server/define/http_error_code_define"
	"lemon-robot-server/dto"
	"lemon-robot-server/entity"
	"lemon-robot-server/service/service_auth"
	"lemon-robot-server/service/service_user"
)

const urlPrefix = "/user"
const urlLogin = "/login"

func RegApis(router *gin.RouterGroup) {
	router.POST(urlPrefix+urlLogin, login)

	router.POST(urlPrefix, create)
}

func login(ctx *gin.Context) {
	reqAuthUser := dto.LrUserAuthReq{}
	http_common.DealError(ctx, ctx.BindJSON(&reqAuthUser), "", func(ctx *gin.Context) {
		result, userEntity := service_user.CheckUser(reqAuthUser.Number, reqAuthUser.Password)
		if result {
			http_common.Success(ctx, service_auth.GenerateJwtTokenStr(userEntity.UserKey))
		} else {
			http_common.Failed(ctx, http_error_code_define.User_LoginIdentityInfoVerifyFailed)
		}
	})
}

func create(ctx *gin.Context) {
	createUserReq := dto.LrUserCreateReq{}
	http_common.DealError(ctx, ctx.Bind(&createUserReq), "", func(ctx *gin.Context) {
		count := dao_user.CountByUserExample(&entity.User{UserNumber: createUserReq.Number})
		fmt.Println(count)
		if count > 0 {
			http_common.Failed(ctx, http_error_code_define.User_CreateFailedNumberAlreadyExists)
			return
		}
		err, _ := service_user.CreateUser(createUserReq.Number, createUserReq.Password)
		if err == nil {
			http_common.Success(ctx, true)
		} else {
			http_common.Failed(ctx, err.Error())
		}
	})
}
