package controller_task

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-server/controller/http_common"
	"lemon-robot-server/entity"
	"lemon-robot-server/service/service_task"
)

const urlPrefix = "/task"

func RegApis(router *gin.RouterGroup) {
	router.POST(urlPrefix, create)
}

func create(ctx *gin.Context) {
	err := service_task.Create(entity.Task{
		TaskName: "哈哈任务",
		BelongNamespace: entity.Namespace{
			NamespaceName: "哈哈命名空间",
			NamespaceTag:  "ttorg",
		},
		TaskTag: "tttask",
	})
	if err != nil {
		logger.Error("Hlee", err)
	}
	http_common.Success(ctx, "ok")
	//createUserReq := model.ReqUserCreate{}
	//http_common.DealError(ctx, ctx.Bind(&createUserReq), "", func(ctx *gin.Context) {
	//	result, _ := service_user.CreateUser(createUserReq.Number, createUserReq.Password)
	//	http_common.Success(ctx, result)
	//})
}
