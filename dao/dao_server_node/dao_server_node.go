package dao_server_node

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
	"lemon-robot-server/sysinfo"
	"time"
)

func Save(entity *entity.ServerNode) *gorm.DB {
	return db.Db().Save(entity)
}

func FirstByExample(example *entity.ServerNode) entity.ServerNode {
	result := entity.ServerNode{}
	db.Db().First(&result, example)
	return result
}

func FindAllByExample(example *entity.ServerNode) []entity.ServerNode {
	result := make([]entity.ServerNode, 3)
	//db.Db().Find(&result, example)
	db.Db().Set("gorm:auto_preload", true).Find(&result, example)
	return result
}

func FindAllActiveNode() []entity.ServerNode {
	result := make([]entity.ServerNode, 3)
	// active容差为activeTime的二倍
	dur, _ := time.ParseDuration(fmt.Sprintf("-%ds", sysinfo.LrServerConfig().ClusterNodeActiveInterval*2))
	db.Db().Where("active_time > ?", time.Now().Add(dur)).Find(&result)
	return result
}

func CountByUserExample(example *entity.ServerNode) int {
	var count int
	db.Db().Model(&entity.ServerNode{}).Where(example).Count(&count)
	return count
}

func UpdateActiveTime(machineSign string, activeTime time.Time) {
	db.Db().Model(&entity.ServerNode{}).Where("machine_sign = ?", machineSign).Update("active_time", activeTime)
}

func UpdateAlias(machineSign, newAlias string) {
	db.Db().Model(&entity.ServerNode{}).Where("machine_sign = ?", machineSign).Update("alias", newAlias)
}
