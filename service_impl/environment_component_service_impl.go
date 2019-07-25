package service_impl

import (
	"lemon-robot-golang-commons/utils/lru_string"
	"lemon-robot-server/dao"
	"lemon-robot-server/entity"
	"time"
)

type EnvironmentComponentServiceImpl struct {
	environmentComponentDao *dao.EnvironmentComponentDao
}

func NewEnvironmentComponentServiceImpl() *EnvironmentComponentServiceImpl {
	return &EnvironmentComponentServiceImpl{
		environmentComponentDao : dao.NewEnvironmentComponentDao(),
	}
}

func (i *EnvironmentComponentServiceImpl) Save(been *entity.EnvironmentComponent) (error, entity.EnvironmentComponent) {
	if been.EnvironmentComponentKey == "" {
		been.EnvironmentComponentKey = lru_string.GetInstance().Uuid(true)
		been.CreatedAt = time.Now()
		return i.environmentComponentDao.Create(been)
	}else {
		error, queryEnvironmentComponent := i.environmentComponentDao.QueryOne(been.EnvironmentComponentKey)
		if error != nil {
			return error, queryEnvironmentComponent
		}
		been.UpdatedAt = time.Now()
		return i.environmentComponentDao.Update(been)
	}
}

func (i *EnvironmentComponentServiceImpl) Delete(key string) error {
	error, queryEnviromentComponent := i.environmentComponentDao.QueryOne(key)
	if error != nil {
		return error
	}
	currentTime:=time.Now()
	//queryEnviromentComponent.DeletedAt = &currentTime
	return i.environmentComponentDao.Delete(&queryEnviromentComponent, currentTime)
}

func (i *EnvironmentComponentServiceImpl) QueryList() (error, []entity.EnvironmentComponent) {
	return i.environmentComponentDao.QueryList()
}