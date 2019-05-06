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
	db.Db().Find(&result, example)
	return result
}

func FindAllExpiredNode() []entity.ServerNode {
	result := make([]entity.ServerNode, 3)
	dur, _ := time.ParseDuration(fmt.Sprintf("-%ds", sysinfo.LrConfig().ClusterNodeActiveInterval))
	db.Db().Where("active_time < ?", time.Now().Add(dur)).Find(&result)
	return result
}

func CountByUserExample(example *entity.ServerNode) int {
	var count int
	db.Db().Model(&entity.ServerNode{}).Where(example).Count(&count)
	return count
}

func UpdateActiveTime(machineCode string, activeTime time.Time) {
	db.Db().Model(&entity.ServerNode{}).Where("machine_Code = ?", machineCode).Update("active_time", activeTime)
}
