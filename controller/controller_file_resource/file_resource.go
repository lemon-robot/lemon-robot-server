package controller_file_resource

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-server/controller/http_response"
	"lemon-robot-server/service/service_file_resource"
)

const urlPrefix = "/file_resource"
const formFileKey = "file"

func RegApis(engine *gin.Engine) {
	engine.POST(urlPrefix, generateFileResourceKey)
	engine.PUT(urlPrefix, uploadFileResource)
}

func generateFileResourceKey(ctx *gin.Context) {
	http_response.Success(ctx, service_file_resource.GenerateFileResourceKey())
}

func uploadFileResource(ctx *gin.Context) {
	file, fileHeader, err := ctx.Request.FormFile(formFileKey)
	if err != nil {
		http_response.Failed(ctx, http_response.ErrCode_FileResource_AnalysisFailed)
	} else {
		fileResourceKey := ctx.Request.Header["File_resource_key"][0]
		success, code := service_file_resource.UploadFileResource(file, fileHeader, fileResourceKey)
		http_response.Response(ctx, success, code, nil)
	}
}
