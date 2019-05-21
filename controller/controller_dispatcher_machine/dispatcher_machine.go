package controller_dispatcher_machine

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-server/controller/http_common"
	"lemon-robot-server/define/http_error_code_define"
	"lemon-robot-server/dto"
	"lemon-robot-server/service"
	"lemon-robot-server/service_impl"
)

var dispatcherMachineService service.DispatcherMachineService = service_impl.NewDispatcherMachineServiceImpl()

const urlPrefix = "/dispatcher_machine"

func RegApis(router *gin.RouterGroup) {
	router.PUT(urlPrefix+"/alias", setAlias)
	router.PUT(urlPrefix+"/tags", setTags)
}

func setAlias(ctx *gin.Context) {
	req := &dto.CommonMachineSetAliasReq{}
	http_common.DealError(ctx, ctx.BindJSON(&req), "", func(ctx *gin.Context) {
		if dispatcherMachineService.SetAlias(req) {
			http_common.Success(ctx, true)
		} else {
			http_common.Failed(ctx, http_error_code_define.Common_ServerInternalError)
		}
	})
}
func setTags(ctx *gin.Context) {
	req := &dto.DispatcherMachineSetTagsReq{}
	http_common.DealError(ctx, ctx.BindJSON(&req), "", func(ctx *gin.Context) {
		if dispatcherMachineService.SetTags(req) {
			http_common.Success(ctx, true)
		} else {
			http_common.Failed(ctx, http_error_code_define.Common_ServerInternalError)
		}
	})
}
