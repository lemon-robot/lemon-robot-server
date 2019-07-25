package service_impl

import (
	"lemon-robot-golang-commons/utils/lru_string"
	"lemon-robot-server/dao"
	"lemon-robot-server/dto"
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

func (i *EnvironmentComponentServiceImpl) Save(environmentComponentReq *dto.EnvironmentComponentReq) (error, dto.EnvironmentComponentReq) {
	environmentComponent := entity.EnvironmentComponent{}
	environmentComponent.EnvironmentComponentKey = environmentComponentReq.EnvironmentComponentKey
	environmentComponent.EnvironmentComponentDescription = environmentComponentReq.EnvironmentComponentDescription
	environmentComponent.EnvironmentComponentName = environmentComponentReq.EnvironmentComponentName
	environmentComponent.CreatedAt = environmentComponentReq.CreatedAt
	environmentComponent.UpdatedAt = environmentComponentReq.UpdatedAt
	if environmentComponentReq.EnvironmentComponentKey == "" {
		environmentComponentReq.EnvironmentComponentKey = lru_string.GetInstance().Uuid(true)
		environmentComponentReq.CreatedAt = time.Now()
		environmentComponent.EnvironmentComponentKey = environmentComponentReq.EnvironmentComponentKey
		environmentComponent.CreatedAt = environmentComponentReq.CreatedAt
		error, _ := i.environmentComponentDao.Create(&environmentComponent)
		return error, *environmentComponentReq
	}else {
		error, _ := i.environmentComponentDao.QueryOne(environmentComponentReq.EnvironmentComponentKey)
		if error != nil {
			return error, *environmentComponentReq
		}
		environmentComponentReq.UpdatedAt = time.Now()
		err, _ := i.environmentComponentDao.Update(&environmentComponent)
		return err, *environmentComponentReq
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

func (i *EnvironmentComponentServiceImpl) QueryList() (error, []dto.EnvironmentComponentReq) {
	var environmentComponentReqs []dto.EnvironmentComponentReq
	error, environmentComponents := i.environmentComponentDao.QueryList()
	for _, v := range environmentComponents {
		environmentComponentReq := dto.EnvironmentComponentReq{}
		environmentComponentReq.EnvironmentComponentKey = v.EnvironmentComponentKey
		environmentComponentReq.EnvironmentComponentName = v.EnvironmentComponentName
		environmentComponentReq.EnvironmentComponentDescription = v.EnvironmentComponentDescription
		environmentComponentReq.CreatedAt = v.CreatedAt
		environmentComponentReq.UpdatedAt = v.UpdatedAt
		environmentComponentReqs = append(environmentComponentReqs, environmentComponentReq)
	}
	return error, environmentComponentReqs
}