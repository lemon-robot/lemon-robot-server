package controller_environment

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lemon-robot-server/controller/http_common"
	"lemon-robot-server/dto"
	"lemon-robot-server/entity"
	"lemon-robot-server/service"
	"lemon-robot-server/service_impl"
)


var environmentService service.EnvironmentService = service_impl.NewEnvironmentComponentServiceImpl()

const urlPrefix = "/environment_component"

func RegApis(router *gin.RouterGroup)  {
	router.PUT(urlPrefix, save)
	router.GET(urlPrefix, queryList)
	router.DELETE(urlPrefix, delete)
}

func save(ctx *gin.Context)  {
	environmentComponent := &dto.EnvironmentComponentReq{}
	err := ctx.BindJSON(&environmentComponent)
	if err != nil {
		fmt.Println("get request params error : ", err)
		return
	}
	error, been := environmentService.Save(environmentComponent)
	if error != nil {
		http_common.Failed(ctx, error.Error())
	}else {
		http_common.Success(ctx, been.EnvironmentComponentKey)
	}
}

func queryList(ctx *gin.Context)  {
	error, environmentComponents := environmentService.QueryList()
	if error != nil {
		http_common.Failed(ctx, error.Error())
	}else {
		http_common.Success(ctx, environmentComponents)
	}
}

func delete(ctx *gin.Context)  {
	environmentComponent := &entity.EnvironmentComponent{}
	err := ctx.BindJSON(&environmentComponent)
	if err != nil {
		fmt.Println("get request params error : ", err)
		return
	}
	error := environmentService.Delete(environmentComponent.EnvironmentComponentKey)
	if error != nil {
		http_common.Failed(ctx, error.Error())
	}else {
		http_common.Success(ctx, "")
	}
}