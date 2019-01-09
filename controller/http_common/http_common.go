package http_common

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-golang-commons/model"
)

const BaseUrlPathPrefixFree = "/free"
const BaseUrlPathPrefixMonigor = "/monitor"

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
