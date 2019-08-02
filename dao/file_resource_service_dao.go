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