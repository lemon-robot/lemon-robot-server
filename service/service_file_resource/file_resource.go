package service_file_resource

import (
	"lemon-robot-golang-commons/utils/lrustring"
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
)

func GenerateFRKey() string {
	fileKey := lrustring.Uuid()
	frObj := entity.FileResource{
		FileKey:     fileKey,
		UploadedTag: true,
		UsedTag:     false,
	}
	db.Db().Create(&frObj)
	return fileKey
}
