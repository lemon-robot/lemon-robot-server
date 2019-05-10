package dao

import (
	"github.com/jinzhu/gorm"
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
)

type DispatcherOnlineDao struct{}

func NewDispatcherOnlineDao() *DispatcherOnlineDao {
	return &DispatcherOnlineDao{}
}

func (i *DispatcherOnlineDao) Delete(query interface{}, where ...interface{}) {
	db.Db().Where(query, where).Delete(&entity.DispatcherOnline{})
}

func (i *DispatcherOnlineDao) DeleteByServerNodeMachineSign(nodeMachineSign string) {
	i.Delete("bind_server_machine_sign = ?", nodeMachineSign)
}

func (i *DispatcherOnlineDao) Save(entity *entity.DispatcherOnline) *gorm.DB {
	return db.Db().Save(entity)
}

func (i *DispatcherOnlineDao) FirstByExample(example *entity.DispatcherOnline) entity.DispatcherOnline {
	result := entity.DispatcherOnline{}
	db.Db().First(&result, example)
	return result
}

func (i *DispatcherOnlineDao) CountByUserExample(example *entity.DispatcherOnline) int {
	var count int
	db.Db().Model(&entity.DispatcherOnline{}).Where(example).Count(&count)
	return count
}
