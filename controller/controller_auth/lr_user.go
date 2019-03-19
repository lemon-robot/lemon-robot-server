package controller_auth

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-server/controller/http_common"
	"lemon-robot-server/model"
	"lemon-robot-server/service/service_auth"
	"lemon-robot-server/service/service_lr_user"
)

func authLrUser(ctx *gin.Context) {
	reqAuthLrc := model.ReqAuthLrUser{}
	http_common.DealError(ctx, ctx.BindJSON(&reqAuthLrc), "", func(ctx *gin.Context) {
		result, userEntity := service_lr_user.CheckLrc(reqAuthLrc.Number, reqAuthLrc.Password)
		if result {
			http_common.Success(ctx, service_auth.GenerateJwtTokenStr(userEntity.LrUserKey))
		} else {
			http_common.Failed(ctx, http_common.ErrCode_Auth_IdentityInfoVerifyFailed)
		}
	})
}

func createLrUser(ctx *gin.Context) {
	reqLrcGenerate := model.ReqLrUserCreate{}
	http_common.DealError(ctx, ctx.Bind(&reqLrcGenerate), "", func(ctx *gin.Context) {
		service_lr_user.GenerateLrUser(reqLrcGenerate.Number, reqLrcGenerate.Password)
		http_common.Success(ctx, true)
	})
}
