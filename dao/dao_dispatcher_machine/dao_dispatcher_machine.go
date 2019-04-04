package dao_dispatcher_machine

import (
	"github.com/jinzhu/gorm"
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
)

func Save(entity *entity.DispatcherMachine) *gorm.DB {
	return db.Db().Save(entity)
}

func FirstByExample(example *entity.DispatcherMachine) entity.DispatcherMachine {
	result := entity.DispatcherMachine{}
	db.Db().First(&result, example)
	return result
}

func CountByUserExample(example *entity.DispatcherMachine) int {
	var count int
	db.Db().Model(&entity.DispatcherMachine{}).Where(example).Count(&count)
	return count
}
