package dao

import (
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
	"time"
)

type EnvironmentComponentDao struct {

}

func NewEnvironmentComponentDao() *EnvironmentComponentDao {
	return &EnvironmentComponentDao{}
}

func (i *EnvironmentComponentDao)Create(been *entity.EnvironmentComponent) (error, entity.EnvironmentComponent) {
	result := db.Db().Create(been)
	return result.Error, *been
}

func (i *EnvironmentComponentDao)Delete(been *entity.EnvironmentComponent, time time.Time) error {
	resutl := db.Db().Model(been).Update("DeletedAt", time)
	return resutl.Error
}

func (i *EnvironmentComponentDao)Update(been *entity.EnvironmentComponent) (error, entity.EnvironmentComponent) {
	result := db.Db().Save(been)
	return result.Error, *been
}

func (i *EnvironmentComponentDao)QueryList() (error, []entity.EnvironmentComponent) {
	var environmentComponents []entity.EnvironmentComponent
	//queryResult := db.Db().Find(&environmentComponents)
	queryResult := db.Db().Order("created_at desc").Find(&environmentComponents)
	return queryResult.Error, environmentComponents
}

func (i *EnvironmentComponentDao)QueryOne(environmentComponentKey string) (error, entity.EnvironmentComponent) {
	environmentComponent := entity.EnvironmentComponent{}
	example := entity.EnvironmentComponent{}
	example.EnvironmentComponentKey = environmentComponentKey
	query := db.Db().First(&environmentComponent, &example)
	return query.Error, environmentComponent
}