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
		http_common.Success(ctx, fileResourceService.Upload(file, handler))
	}
}

//func create(ctx *gin.Context) {
//	http_common.Success(ctx, service_file_resource.GenerateFileResourceKey())
//}
//
//func upload(ctx *gin.Context) {
//	file, fileHeader, err := ctx.Request.FormFile(formFileKey)
//	if err != nil {
//		http_common.Failed(ctx, http_error_code_define.FileResource_AnalysisFailed)
//	} else {
//		fileResourceKey := ctx.Request.Header[headerFileResourceKeyField][0]
//		success, code := service_file_resource.UploadFileResource(file, fileHeader, fileResourceKey)
//		http_common.Response(ctx, success, code, nil)
//	}
//}
