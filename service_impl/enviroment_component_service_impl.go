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
		error, queryEnviromentComponent := i.enviromentComponentDao.QueryOne(been.EnvironmentComponentKey)
		if error != nil {
			return error, queryEnviromentComponent
		}
		been.UpdatedAt = time.Now()
		return i.enviromentComponentDao.Update(been)
	}
}

func (i *EnviromentComponentServiceImpl) Delete(key string) error {
	error, queryEnviromentComponent := i.enviromentComponentDao.QueryOne(key)
	if error != nil {
		return error
	}
	currentTime:=time.Now()
	queryEnviromentComponent.DeletedAt = &currentTime
	return i.enviromentComponentDao.Delete(&queryEnviromentComponent, currentTime)
}

func (i *EnviromentComponentServiceImpl) QueryList() (error, []entity.EnvironmentComponent) {
	return i.enviromentComponentDao.QueryList()
}