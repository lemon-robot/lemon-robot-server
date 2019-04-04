package dao_namespace

import (
	"github.com/jinzhu/gorm"
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
)

func Save(entity *entity.Namespace) *gorm.DB {
	return db.Db().Save(entity)
}

func FirstByExample(example *entity.Namespace) entity.Namespace {
	result := entity.Namespace{}
	db.Db().First(&result, example)
	return result
}

func CountByUserExample(example *entity.Namespace) int {
	var count int
	db.Db().Model(&entity.Namespace{}).Where(example).Count(&count)
	return count
}
