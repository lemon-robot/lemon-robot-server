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
	// 先在磁盘中接收文件,然后再上传到oss中, 完成后再返回给客户端key(uuid), 最后删除磁盘文件
	file, handler, err := ctx.Request.FormFile("file")
	if err != nil {
		fmt.Println("err === ", err.Error())
		http_common.Failed(ctx, err.Error())
	}else {
		result, error := fileResourceService.Upload(file, handler)
		if error != nil {
			http_common.Failed(ctx, error.Error())
		}else {
			http_common.Success(ctx, result)
		}
	}
}

