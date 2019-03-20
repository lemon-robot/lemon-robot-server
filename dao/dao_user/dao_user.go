package dao_user

import (
	"github.com/jinzhu/gorm"
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
)

func Save(entity *entity.User) *gorm.DB {
	return db.Db().Save(entity)
}

func FirstByExample(example *entity.User) entity.User {
	result := entity.User{}
	db.Db().First(&result, example)
	return result
}

func CountByUserExample(example *entity.User) int {
	var count int
	db.Db().Model(&entity.User{}).Where(example).Count(&count)
	return count
}
