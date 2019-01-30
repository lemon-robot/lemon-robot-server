package controller_auth

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-server/controller/http_common"
	"lemon-robot-server/model"
	"lemon-robot-server/service/service_auth"
	"lemon-robot-server/service/service_lrc"
)

func authLrc(ctx *gin.Context) {
	reqAuthLrc := model.ReqAuthLrc{}
	http_common.DealError(ctx, ctx.BindJSON(&reqAuthLrc), "", func(ctx *gin.Context) {
		if service_lrc.CheckLrc(reqAuthLrc.Lrct, reqAuthLrc.Lrcp) {
			http_common.Success(ctx, service_auth.GenerateJwtTokenStr(reqAuthLrc.Lrct))
		} else {
			http_common.Failed(ctx, http_common.ErrCode_Auth_IdentityInfoVerifyFailed)
		}
	})
}

func generateLrc(ctx *gin.Context) {
	reqLrcGenerate := model.ReqLrcGenerate{}
	http_common.DealError(ctx, ctx.Bind(&reqLrcGenerate), "", func(ctx *gin.Context) {
		lrcResult := model.ReqAuthLrc{
			Lrct: service_lrc.GenerateLrc(reqLrcGenerate.Lrcp).Lrct,
			Lrcp: reqLrcGenerate.Lrcp,
		}
		http_common.Success(ctx, lrcResult)
	})
}
