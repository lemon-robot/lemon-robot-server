package dao

import (
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
)

type OperatePlatformDao struct {

}

func NewOperatePlatformDao() *OperatePlatformDao {
	return &OperatePlatformDao{}
}

func (i *OperatePlatformDao) GetAll() ([]entity.OperatePlatform, error) {
	var operatePlatforms []entity.OperatePlatform
	result := db.Db().Order("created_at desc").Find(&operatePlatforms)
	return operatePlatforms, result.Error
}

func (i *OperatePlatformDao) GetOnes(key string) (entity.OperatePlatform, error) {
	operatePlatform := entity.OperatePlatform{}
	example := entity.OperatePlatform{}
	example.OperatePlatformKey = key
	result := db.Db().First(&operatePlatform, example)
	return operatePlatform, result.Error
}