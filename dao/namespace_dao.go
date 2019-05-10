package dao

import (
	"github.com/jinzhu/gorm"
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
)

type NameSpaceDao struct{}

func NewNameSpaceDao() *NameSpaceDao {
	return &NameSpaceDao{}
}

func (i *NameSpaceDao) Save(entity *entity.Namespace) *gorm.DB {
	return db.Db().Save(entity)
}

func (i *NameSpaceDao) FirstByExample(example *entity.Namespace) entity.Namespace {
	result := entity.Namespace{}
	db.Db().First(&result, example)
	return result
}

func (i *NameSpaceDao) CountByUserExample(example *entity.Namespace) int {
	var count int
	db.Db().Model(&entity.Namespace{}).Where(example).Count(&count)
	return count
}
