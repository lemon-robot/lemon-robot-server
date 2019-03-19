package dao_user

import (
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
)

func FirstByExample(entityExample entity.User) entity.User {
	user := entity.User{}
	db.Db().First(&user, entityExample)
	return user
}
