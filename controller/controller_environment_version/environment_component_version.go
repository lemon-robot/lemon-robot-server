package controller_environment_version

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-server/controller/http_common"
	"lemon-robot-server/dto"
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
	environmentComponentVersion := &dto.EnvironmentComponentVersionReq{}
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
	belongEnvironmentComponentKey := ctx.Query("belongEnvironmentComponentKey")
	logger.Info("environment_component_version interface queryList:  belongEnvironmentComponentKey == " + belongEnvironmentComponentKey)
	error, environmentComponentVersions := environmentComponentVersionService.QueryList(belongEnvironmentComponentKey)
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
	err := environmentComponentVersionService.Delete(environmentComponentVersion.ECVersionKey)
	if err != nil {
		http_common.Failed(ctx, err.Error())
	}else {
		http_common.Success(ctx, "")
	}
}