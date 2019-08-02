package controller_operate_platform

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-server/controller/http_common"
	"lemon-robot-server/service"
	"lemon-robot-server/service_impl"
)

const urlPrefix = "/operate_platform"

var operatePlatformService service.OperatePlatformService = service_impl.NewOperatePlatformServiceImpl()

func RegApis(router *gin.RouterGroup) {
	router.GET(urlPrefix, getAll)
}

func getAll(ctx *gin.Context) {
	operatePlatforms, error := operatePlatformService.GetAll()
	if error != nil {
		http_common.Failed(ctx, error.Error())
	}else {
		http_common.Success(ctx, operatePlatforms)
	}
}
