package dao_dispatcher_online

import (
	"github.com/jinzhu/gorm"
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
)

type DispatcherOnlineDao struct{}

func (dod DispatcherOnlineDao) Delete(query interface{}, where ...interface{}) {
	db.Db().Where(query, where).Delete(&entity.DispatcherOnline{})
}

func (dod DispatcherOnlineDao) DeleteByClusterNodeMachineSign(nodeMachineSign string) {
	dod.Delete("bind_server_machine_sign = ?", nodeMachineSign)
}

func (dod DispatcherOnlineDao) Save(entity *entity.DispatcherOnline) *gorm.DB {
	return db.Db().Save(entity)
}

func (dod DispatcherOnlineDao) FirstByExample(example *entity.DispatcherOnline) entity.DispatcherOnline {
	result := entity.DispatcherOnline{}
	db.Db().First(&result, example)
	return result
}

func (dod DispatcherOnlineDao) CountByUserExample(example *entity.DispatcherOnline) int {
	var count int
	db.Db().Model(&entity.DispatcherOnline{}).Where(example).Count(&count)
	return count
}
