package controller_file_resource

import (
	"github.com/gin-gonic/gin"
	"lemon-robot-server/controller/http_response"
	"lemon-robot-server/service/service_file_resource"
)

const frFolderName string = "file_resource"
const urlPrefix = "/file_resource"

func RegApis(engine *gin.Engine) {
	engine.POST(urlPrefix, generateFRKey)
}

func generateFRKey(ctx *gin.Context) {
	http_response.Success(ctx, service_file_resource.GenerateFRKey())
}

func upload(ctx *gin.Context) {
	//uploadFile, _, _ := ctx.Request.FormFile("upload")
	//io.CopyFile(uploadFile, uploadFile)
}
