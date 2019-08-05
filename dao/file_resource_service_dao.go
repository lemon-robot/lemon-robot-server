package dao

import (
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
)

type FileResourceServiceDao struct {
}

func NewFileResourceServiceDao() *FileResourceServiceDao {
	return &FileResourceServiceDao{}
}

func (i *FileResourceServiceDao) Create(been *entity.FileResource) error {
	result := db.Db().Create(been)
	return result.Error
}

func (i *FileResourceServiceDao) GetFileSource(fileSourceKey string) (entity.FileResource, error) {
	var example entity.FileResource
	example.FileResourceKey = fileSourceKey
	fileSource := entity.FileResource{}
	resutl := db.Db().First(&fileSource, &example)
	return fileSource, resutl.Error
}

