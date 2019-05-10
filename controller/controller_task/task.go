package controller_task

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-server/controller/http_common"
	"lemon-robot-server/entity"
	"lemon-robot-server/service"
	"lemon-robot-server/service_impl"
)

var taskService service.TaskService = service_impl.NewTaskServiceImpl()

const urlPrefix = "/task"

func RegApis(router *gin.RouterGroup) {
	router.POST(urlPrefix, create)
}

func create(ctx *gin.Context) {
	err := taskService.Create(entity.Task{
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
	//	result, _ := service_user.Create(createUserReq.Number, createUserReq.Password)
	//	http_common.Success(ctx, result)
	//})
}
