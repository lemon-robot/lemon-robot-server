package service

import "lemon-robot-server/entity"

type DispatcherOnlineService interface {
	DeleteByOnlineKey(onlineKey string)
	Save(dispatcherOnline *entity.DispatcherOnline)
}
