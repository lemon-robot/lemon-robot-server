package controller_environment_version

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lemon-robot-server/controller/http_common"
	"lemon-robot-server/entity"
	"lemon-robot-server/service"
	"lemon-robot-server/service_impl"
)

var environmentComponentVersionService service.EnvironmentComponentVersionService = service_impl.NewEnvironmentComponentVersionServiceImpl()

const urlPrefix = "/environment_component_version"

func RegApis(router *gin.RouterGroup)()  {
	router.PUT(urlPrefix, save)
	router.GET(urlPrefix, queryList)
	router.DELETE(urlPrefix, delete)
}

func save(ctx *gin.Context)  {
	environmentComponentVersion := &entity.EnvironmentComponentVersion{}
	error := ctx.BindJSON(&environmentComponentVersion)
	if error != nil {
		fmt.Println("get request params error : ", error)
		return
	}
	err, been := environmentComponentVersionService.Save(environmentComponentVersion)
	if err != nil {
		http_common.Failed(ctx, err.Error())
	}else {
		http_common.Success(ctx, been.BelongEnvironmentComponentKey)
	}
}

func queryList(ctx *gin.Context)  {
	error, environmentComponentVersions := environmentComponentVersionService.QueryList("sss")
	if error != nil {
		http_common.Failed(ctx, error.Error())
	}else {
		http_common.Success(ctx, environmentComponentVersions)
	}
}

func delete(ctx *gin.Context)  {
	environmentComponentVersion := &entity.EnvironmentComponentVersion{}
	error := ctx.BindJSON(environmentComponentVersion)
	if error != nil {
		fmt.Println("get request params error : ", error)
		return
	}
	err := environmentComponentVersionService.Delete(environmentComponentVersion.BelongEnvironmentComponentKey)
	if err != nil {
		http_common.Failed(ctx, err.Error())
	}else {
		http_common.Success(ctx, "")
	}
}