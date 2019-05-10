package service_impl

import (
	"lemon-robot-server/dao"
	"lemon-robot-server/entity"
)

type DispatcherOnlineServiceImpl struct {
	dispatcherOnlineDao *dao.DispatcherOnlineDao
}

func NewDispatcherOnlineServiceImpl() *DispatcherOnlineServiceImpl {
	return &DispatcherOnlineServiceImpl{
		dispatcherOnlineDao: dao.NewDispatcherOnlineDao(),
	}
}

func (i *DispatcherOnlineServiceImpl) DeleteByOnlineKey(onlineKey string) {
	i.dispatcherOnlineDao.Delete("online_key", onlineKey)
}

func (i *DispatcherOnlineServiceImpl) Save(dispatcherOnline *entity.DispatcherOnline) {
	i.dispatcherOnlineDao.Save(dispatcherOnline)
}
