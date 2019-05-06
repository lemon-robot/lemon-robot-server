package dao_dispatcher_online

import (
	"github.com/jinzhu/gorm"
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
)

func Delete(query interface{}, where ...interface{}) {
	db.Db().Where(query, where).Delete(&entity.DispatcherOnline{})
}

func DeleteByClusterNodeMachineSign(nodeMachineSign string) {
	Delete("bind_server_machine_sign = ?", nodeMachineSign)
}

func Save(entity *entity.DispatcherOnline) *gorm.DB {
	return db.Db().Save(entity)
}

func FirstByExample(example *entity.DispatcherOnline) entity.DispatcherOnline {
	result := entity.DispatcherOnline{}
	db.Db().First(&result, example)
	return result
}

func CountByUserExample(example *entity.DispatcherOnline) int {
	var count int
	db.Db().Model(&entity.DispatcherOnline{}).Where(example).Count(&count)
	return count
}
