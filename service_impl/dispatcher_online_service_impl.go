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
	i.dispatcherOnlineDao.Delete(&entity.DispatcherOnline{
		OnlineKey: onlineKey,
	})
}

func (i *DispatcherOnlineServiceImpl) Save(dispatcherOnline *entity.DispatcherOnline) {
	i.dispatcherOnlineDao.Save(dispatcherOnline)
}

func (i *DispatcherOnlineServiceImpl) ClearAllOffline() {
	i.dispatcherOnlineDao.ClearAllOffline()
}
