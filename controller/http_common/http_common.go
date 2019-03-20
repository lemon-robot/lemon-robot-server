package http_common

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/model"
	"lemon-robot-server/define/http_error_code_define"
)

func Success(ctx *gin.Context, data interface{}) {
	Response(ctx, true, "", data)
}

func Failed(ctx *gin.Context, code string) {
	Response(ctx, false, code, nil)
}

func Response(ctx *gin.Context, success bool, code string, data interface{}) {
	ctx.JSON(200, model.HttpResponse{
		Success: success,
		Code:    code,
		Data:    data,
	})
}

func DealError(ctx *gin.Context, err error, tip string, noErrCallback func(ctx *gin.Context)) {
	if err != nil {
		if tip == "" {
			tip = "An error has occurred inside the system"
		}
		logger.Error(tip, err)
		Failed(ctx, http_error_code_define.Common_ServerInternalError)
	} else {
		noErrCallback(ctx)
	}
}
