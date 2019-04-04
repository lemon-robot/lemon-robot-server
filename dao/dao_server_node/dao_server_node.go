package dao_server_node

import (
	"github.com/jinzhu/gorm"
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
)

func Save(entity *entity.ServerNode) *gorm.DB {
	return db.Db().Save(entity)
}

func FirstByExample(example *entity.ServerNode) entity.ServerNode {
	result := entity.ServerNode{}
	db.Db().First(&result, example)
	return result
}

func CountByUserExample(example *entity.ServerNode) int {
	var count int
	db.Db().Model(&entity.ServerNode{}).Where(example).Count(&count)
	return count
}
