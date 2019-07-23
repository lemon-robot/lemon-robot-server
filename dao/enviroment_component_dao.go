package dao

import (
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
	"time"
)

type EnviromentComponentDao struct {

}

func NewEnviromentComponentDao() *EnviromentComponentDao {
	return &EnviromentComponentDao{}
}

func (i *EnviromentComponentDao)Create(been *entity.EnvironmentComponent) (error, entity.EnvironmentComponent) {
	//return db.Db().Save(entity)
	result := db.Db().Create(been)
	return result.Error, *been
}

func (i *EnviromentComponentDao)Delete(been *entity.EnvironmentComponent, time time.Time) error {
	resutl := db.Db().Model(been).Update("DeletedAt", time)
	return resutl.Error
}

func (i *EnviromentComponentDao)Update(been *entity.EnvironmentComponent) (error, entity.EnvironmentComponent) {
	result := db.Db().Save(been)
	return result.Error, *been
}

func (i *EnviromentComponentDao)Query(environmentComponentKey string) (error, entity.EnvironmentComponent) {
	result := entity.EnvironmentComponent{}
	queryResult := db.Db().First(&result, environmentComponentKey)
	return queryResult.Error, result
}