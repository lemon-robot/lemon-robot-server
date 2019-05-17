package controller_dispatcher_tag

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-server/controller/http_common"
	"lemon-robot-server/dto"
	"lemon-robot-server/service"
	"lemon-robot-server/service_impl"
)

var dispatcherTagService service.DispatcherTagService = service_impl.NewDispatcherTagServiceImpl()

const urlPrefix = "/dispatcher_tag"

func RegApis(router *gin.RouterGroup) {
	router.GET(urlPrefix+"/list", list)
	router.PUT(urlPrefix+"/save", save)
	router.PUT(urlPrefix+"/save/:tagKey", deleteByTagKey)
}

func list(ctx *gin.Context) {
	http_common.Success(ctx, dispatcherTagService.List())
}

func save(ctx *gin.Context) {
	tagSaveReq := &dto.DispatcherTagSaveReq{}
	http_common.DealError(ctx, ctx.BindJSON(&tagSaveReq), "", func(ctx *gin.Context) {
		dispatcherTagService.Save(tagSaveReq)
		http_common.Success(ctx, true)
	})
}

func deleteByTagKey(ctx *gin.Context) {
	tagKey := ctx.Param("tagKey")
	dispatcherTagService.Delete(tagKey)
	http_common.Success(ctx, true)
}
