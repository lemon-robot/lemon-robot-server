package dao_common

import (
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
)

func FirstByExample(entity, entityExample entity.LrEntity) entity.LrEntity {
	db.Db().First(&entity, entityExample)
	return entity
}
