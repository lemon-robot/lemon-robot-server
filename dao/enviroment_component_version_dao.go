package dao

import (
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
	"time"
)

type EnvironmentComponentVersionDao struct {

}

func NewEnvironmentComponentVersionDao() *EnvironmentComponentVersionDao {
	return &EnvironmentComponentVersionDao{}
}

func (i *EnvironmentComponentVersionDao) Create(been *entity.EnvironmentComponentVersion) (error, entity.EnvironmentComponentVersion) {
	result := db.Db().Create(been)
	return result.Error, *been
}

func (i *EnvironmentComponentVersionDao) Updata(been *entity.EnvironmentComponentVersion) (error, entity.EnvironmentComponentVersion) {
	result := db.Db().Save(been)
	return result.Error, *been
}

func (i *EnvironmentComponentVersionDao) Delete(been *entity.EnvironmentComponentVersion, currentTime time.Time) error {
	result := db.Db().Model(been).Update("DeletedAt", currentTime)
	return result.Error
}

func (i *EnvironmentComponentVersionDao) QueryList(belongEnvironmentComponentKey string) (error, []entity.EnvironmentComponentVersion) {
	var environmentComponentVersions []entity.EnvironmentComponentVersion
	exampleEnvironmentComponentVersion := entity.EnvironmentComponentVersion{}
	exampleEnvironmentComponentVersion.BelongEnvironmentComponentKey = belongEnvironmentComponentKey
	result := db.Db().Find(&environmentComponentVersions, exampleEnvironmentComponentVersion)
	return result.Error, environmentComponentVersions
}

func (i *EnvironmentComponentVersionDao) QueryOne(belongEnvironmentComponentKey string) (error, entity.EnvironmentComponentVersion) {
	resultEnvironmentComponentVersion := entity.EnvironmentComponentVersion{}
	example := entity.EnvironmentComponentVersion{}
	example.BelongEnvironmentComponentKey = belongEnvironmentComponentKey
	result := db.Db().First(&resultEnvironmentComponentVersion, &example)
	return result.Error, resultEnvironmentComponentVersion
}