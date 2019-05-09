package service_timer

import (
	"lemon-robot-server/core/core_other"
	"lemon-robot-server/service/service_server_node"
	"lemon-robot-server/sysinfo"
	"time"
)

func startActiveTimer() {
	core_other.WorkLock.Add(1)
	ticker := time.NewTicker(time.Second * time.Duration(sysinfo.LrServerConfig().ClusterNodeActiveInterval))
	go func() {
		for range ticker.C {
			service_server_node.RefreshActiveTime()
		}
	}()
}
