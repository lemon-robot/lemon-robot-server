package dao

import (
	"fmt"
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

func (i *EnviromentComponentDao)QueryList() (error, []entity.EnvironmentComponent) {
	var environmentComponents []entity.EnvironmentComponent
	queryResult := db.Db().Find(&environmentComponents)
	return queryResult.Error, environmentComponents
}

func (i *EnviromentComponentDao)QueryOne(environmentComponentKey string) (error, entity.EnvironmentComponent) {
	environmentComponent := entity.EnvironmentComponent{}
	example := entity.EnvironmentComponent{}
	example.EnvironmentComponentKey = environmentComponentKey
	query := db.Db().First(&environmentComponent, &example)
	//query := db.Db().Where("environmentComponentKey = ?", environmentComponent).First(environmentComponent)
	if query.Error != nil {
		fmt.Println("query one error ================================== ", query.Error)
	}
	return query.Error, environmentComponent
}