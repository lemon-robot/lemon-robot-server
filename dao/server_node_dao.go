package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
	"lemon-robot-server/sysinfo"
	"time"
)

type ServerNodeDao struct{}

func NewServerNodeDao() *ServerNodeDao {
	return &ServerNodeDao{}
}

func (i *ServerNodeDao) Save(entity *entity.ServerNode) *gorm.DB {
	return db.Db().Save(entity)
}

func (i *ServerNodeDao) FirstByExample(example *entity.ServerNode) entity.ServerNode {
	result := entity.ServerNode{}
	db.Db().First(&result, example)
	return result
}

func (i *ServerNodeDao) FindAllByExample(example *entity.ServerNode) []entity.ServerNode {
	result := make([]entity.ServerNode, 3)
	//db.Db().Find(&result, example)
	db.Db().Set("gorm:auto_preload", true).Find(&result, example)
	return result
}

func (i *ServerNodeDao) FindAllActiveNode() []entity.ServerNode {
	result := make([]entity.ServerNode, 3)
	// active容差为activeTime的二倍
	dur, _ := time.ParseDuration(fmt.Sprintf("-%ds", sysinfo.LrServerConfig().ClusterNodeActiveInterval*2))
	db.Db().Where("active_time > ?", time.Now().Add(dur)).Find(&result)
	return result
}

func (i *ServerNodeDao) CountByUserExample(example *entity.ServerNode) int {
	var count int
	db.Db().Model(&entity.ServerNode{}).Where(example).Count(&count)
	return count
}

func (i *ServerNodeDao) UpdateActiveTime(machineSign string, activeTime time.Time) {
	db.Db().Model(&entity.ServerNode{}).Where("machine_sign = ?", machineSign).Update("active_time", activeTime)
}

func (i *ServerNodeDao) UpdateAlias(machineSign, newAlias string) {
	db.Db().Model(&entity.ServerNode{}).Where("machine_sign = ?", machineSign).Update("alias", newAlias)
}
