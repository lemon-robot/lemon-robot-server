package controller_file_resource

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lemon-robot-server/controller/http_common"
	"lemon-robot-server/service"
	"lemon-robot-server/service_impl"
)

const urlPrefix = "/file_resource"
const formFileKey = "file"
const headerFileResourceKeyField = "File_resource_key"

var fileResourceService service.FileResourceService = service_impl.NewFileResourceServiceImpl()

func RegApis(router *gin.RouterGroup) {
	//router.POST(urlPrefix, create)
	//router.PUT(urlPrefix, upload)
	router.POST(urlPrefix, upload)
}

func upload(ctx *gin.Context)  {
	file, handler, err := ctx.Request.FormFile("file")
	if err != nil {
		fmt.Println("err === ", err.Error())
		http_common.Failed(ctx, err.Error())
	}else {
		result, error := fileResourceService.Upload(ctx, file, handler)
		if error != nil {
			http_common.Failed(ctx, error.Error())
		}else {
			http_common.Success(ctx, result)
		}
	}
}

