package dao

import (
	"github.com/jinzhu/gorm"
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
)

type UserDao struct{}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (i *UserDao) Save(entity *entity.User) *gorm.DB {
	return db.Db().Save(entity)
}

func (i *UserDao) FirstByExample(example *entity.User) entity.User {
	result := entity.User{}
	db.Db().First(&result, example)
	return result
}

func (i *UserDao) CountByUserExample(example *entity.User) int {
	var count int
	db.Db().Model(&entity.User{}).Where(example).Count(&count)
	return count
}
