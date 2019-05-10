package controller_server_node

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-server/controller/http_common"
	"lemon-robot-server/dto"
	"lemon-robot-server/service"
	"lemon-robot-server/service_impl"
)

var serverNodeService service.ServerNodeService = service_impl.NewServerNodeServiceImpl()

const urlPrefix = "/server_node"

func RegApis(router *gin.RouterGroup) {
	router.GET(urlPrefix+"/list", list)
	router.PUT(urlPrefix+"/alias", updateAlias)
}

func list(ctx *gin.Context) {
	http_common.Success(ctx, serverNodeService.QueryAllNodeInfo())
}

func updateAlias(ctx *gin.Context) {
	updateAliasInfo := &dto.ServerNodeUpdateAliasReq{}
	http_common.DealError(ctx, ctx.BindJSON(&updateAliasInfo), "", func(ctx *gin.Context) {
		serverNodeService.UpdateAlias(updateAliasInfo.MachineSign, updateAliasInfo.Alias)
		http_common.Success(ctx, true)
	})
}
