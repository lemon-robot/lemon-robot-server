package service_file_resource

import (
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/utils/lruio"
	"lemon-robot-golang-commons/utils/lrustring"
	"lemon-robot-server/controller/http_response"
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
	"lemon-robot-server/sysinfo"
	"mime/multipart"
	"path"
)

const fileResourceFolderPrefix string = "file_resource/"

func GenerateFileResourceKey() string {
	fileKey := lrustring.Uuid()
	frObj := entity.FileResource{
		FileResourceKey: fileKey,
		UploadedTag:     true,
		UsedTag:         false,
	}
	db.Db().Create(&frObj)
	return fileKey
}

func UploadFileResource(file multipart.File, fileHeader *multipart.FileHeader, fileResourceKey string) (success bool, errCode string) {
	fileResourceObj := entity.FileResource{}
	db.Db().First(&fileResourceObj, &entity.FileResource{FileResourceKey: fileResourceKey})
	if fileResourceObj.ID == 0 {
		return false, http_response.ErrCode_FileResource_KeyInvalid
	}
	fileResourceObj.FileExtension = path.Ext(fileHeader.Filename)
	fileResourceObj.OriginalFileName = fileHeader.Filename
	fileResourceObj.FileSize = fileHeader.Size
	fileResourceObj.ContentType = fileHeader.Header["Content-Type"][0]
	fileResourceObj.UploadedTag = true
	db.Db().Save(&fileResourceObj)
	copyErr := lruio.CopyFileFromReader(file, sysinfo.GetWorkspaceSubPath(fileResourceFolderPrefix+fileResourceKey))
	if copyErr != nil {
		logger.Error("Handling user upload file_resource, error occurs when copying file", copyErr)
		return false, http_response.ErrCode_Common_ServerInternalError
	}
	return true, ""
}
