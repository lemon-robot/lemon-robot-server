package controller_task

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-server/controller/http_common"
	"lemon-robot-server/model"
	"lemon-robot-server/service/service_user"
)

const urlPrefix = "/task"

func RegApis(router *gin.RouterGroup) {
	router.POST(urlPrefix, create)
}

func create(ctx *gin.Context) {
	createUserReq := model.ReqUserCreate{}
	http_common.DealError(ctx, ctx.Bind(&createUserReq), "", func(ctx *gin.Context) {
		result, _ := service_user.CreateUser(createUserReq.Number, createUserReq.Password)
		http_common.Success(ctx, result)
	})
}
