package service_impl

import (
	"lemon-robot-server/service"
	"lemon-robot-server/sysinfo"
	"sync"
	"time"
)

var timerWorkLock sync.WaitGroup

type TimerServiceImpl struct {
	serverNodeService       service.ServerNodeService
	dispatcherOnlineService service.DispatcherOnlineService
}

func NewTimerServiceImpl() *TimerServiceImpl {
	return &TimerServiceImpl{
		serverNodeService:       NewServerNodeServiceImpl(),
		dispatcherOnlineService: NewDispatcherOnlineServiceImpl(),
	}
}

func (i *TimerServiceImpl) StartTimer() {
	i.startActiveTimer()
}

func (i *TimerServiceImpl) startActiveTimer() {
	timerWorkLock.Add(1)
	ticker := time.NewTicker(time.Second * time.Duration(sysinfo.LrServerConfig().ClusterNodeActiveInterval))
	go func() {
		for range ticker.C {
			i.serverNodeService.RefreshActiveTime()
			i.dispatcherOnlineService.ClearAllOffline()
		}
	}()
}
