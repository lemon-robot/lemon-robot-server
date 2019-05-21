package dao

import (
	"github.com/jinzhu/gorm"
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
)

type DispatcherMachineDao struct{}

func NewDispatcherMachineDao() *DispatcherMachineDao {
	return &DispatcherMachineDao{}
}

func (i *DispatcherMachineDao) Save(entity *entity.DispatcherMachine) *gorm.DB {
	return db.Db().Save(entity)
}

func (i *DispatcherMachineDao) FirstByExample(example *entity.DispatcherMachine) entity.DispatcherMachine {
	result := entity.DispatcherMachine{}
	db.Db().First(&result, example)
	return result
}

func (i *DispatcherMachineDao) CountByUserExample(example *entity.DispatcherMachine) int {
	var count int
	db.Db().Model(&entity.DispatcherMachine{}).Where(example).Count(&count)
	return count
}
