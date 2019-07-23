package service_impl

import (
	"lemon-robot-golang-commons/utils/lru_string"
	"lemon-robot-server/dao"
	"lemon-robot-server/entity"
	"time"
)

type EnviromentComponentServiceImpl struct {
	enviromentComponentDao *dao.EnviromentComponentDao
}

func NewEnviromentComponentServiceImpl() *EnviromentComponentServiceImpl {
	return &EnviromentComponentServiceImpl{
		enviromentComponentDao : dao.NewEnviromentComponentDao(),
	}
}

func (i *EnviromentComponentServiceImpl) Save(been *entity.EnvironmentComponent) (error, entity.EnvironmentComponent) {
	if been.EnvironmentComponentKey == "" {
		been.EnvironmentComponentKey = lru_string.GetInstance().Uuid(true)
		been.CreatedAt = time.Now()
		return i.enviromentComponentDao.Create(been)
	}else {
		been.UpdatedAt = time.Now()
		return i.enviromentComponentDao.Update(been)
	}
}

func (i *EnviromentComponentServiceImpl) Delete(key string) error {
	currentTime:=time.Now()
	environmentComponent := entity.EnvironmentComponent{}
	environmentComponent.EnvironmentComponentKey = key
	return i.enviromentComponentDao.Delete(&environmentComponent, currentTime)
}

func (i *EnviromentComponentServiceImpl) Query(key string) (error, entity.EnvironmentComponent) {
	been := &entity.EnvironmentComponent{}
	been.EnvironmentComponentKey = key
	return i.enviromentComponentDao.Query(been.EnvironmentComponentKey)
}