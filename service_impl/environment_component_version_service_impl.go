package service_impl

import (
	"lemon-robot-golang-commons/utils/lru_string"
	"lemon-robot-server/dao"
	"lemon-robot-server/entity"
	"time"
)

type EnvironmentComponentVersionServiceImpl struct {
	environmentComponentVersionDao *dao.EnvironmentComponentVersionDao
}

func NewEnvironmentComponentVersionServiceImpl() *EnvironmentComponentVersionServiceImpl {
	return &EnvironmentComponentVersionServiceImpl{
		environmentComponentVersionDao : dao.NewEnvironmentComponentVersionDao(),
	}
}

func (i *EnvironmentComponentVersionServiceImpl) Save(been *entity.EnvironmentComponentVersion) (error, entity.EnvironmentComponentVersion) {
	ecVersionKey := been.ECVersionKey
	if ecVersionKey == "" {
		been.ECVersionKey = lru_string.GetInstance().Uuid(true)
		been.CreatedAt = time.Now()
		return i.environmentComponentVersionDao.Create(been)
	}else {
		error, queryEnvironmentComponentVersion := i.environmentComponentVersionDao.QueryOne(ecVersionKey)
		if error != nil {
			return error, queryEnvironmentComponentVersion
		}
		been.UpdatedAt = time.Now()
		return i.environmentComponentVersionDao.Updata(been)
	}
}

func (i *EnvironmentComponentVersionServiceImpl) Delete(belongEnvironmentComponentKey string) error {
	error, queryOneiEnvironmentComponentVersion := i.environmentComponentVersionDao.QueryOne(belongEnvironmentComponentKey)
	if error != nil {
		return error
	}
	return i.environmentComponentVersionDao.Delete(&queryOneiEnvironmentComponentVersion, time.Now())
}

func (i *EnvironmentComponentVersionServiceImpl) QueryList(environmentComponentKey string) (error, []entity.EnvironmentComponentVersion) {
	return i.environmentComponentVersionDao.QueryList()
}