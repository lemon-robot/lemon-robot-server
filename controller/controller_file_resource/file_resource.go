package controller_file_resource

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-server/controller/http_common"
	"lemon-robot-server/define/http_error_code_define"
	"lemon-robot-server/service/service_file_resource"
)

const urlPrefix = "/file_resource"
const formFileKey = "file"
const headerFileResourceKeyField = "File_resource_key"

func RegApis(router *gin.RouterGroup) {
	router.POST(urlPrefix, create)
	router.PUT(urlPrefix, upload)
}

func create(ctx *gin.Context) {
	http_common.Success(ctx, service_file_resource.GenerateFileResourceKey())
}

func upload(ctx *gin.Context) {
	file, fileHeader, err := ctx.Request.FormFile(formFileKey)
	if err != nil {
		http_common.Failed(ctx, http_error_code_define.FileResource_AnalysisFailed)
	} else {
		fileResourceKey := ctx.Request.Header[headerFileResourceKeyField][0]
		success, code := service_file_resource.UploadFileResource(file, fileHeader, fileResourceKey)
		http_common.Response(ctx, success, code, nil)
	}
}
