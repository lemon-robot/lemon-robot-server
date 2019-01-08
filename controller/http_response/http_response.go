package http_response

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-golang-commons/model"
)

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, model.HttpResponse{
		Success: true,
		Code:    "",
		Data:    data,
	})
}

func Failed(ctx *gin.Context, code string) {
	ctx.JSON(200, model.HttpResponse{
		Success: false,
		Code:    code,
		Data:    nil,
	})
}
