package service_timer

import (
	"lemon-robot-server/core/core_other"
	"lemon-robot-server/sysinfo"
	"time"
)

func startScanTimer() {
	core_other.WorkLock.Add(1)
	ticker := time.NewTicker(time.Second * time.Duration(sysinfo.LrConfig().ClusterNodeActiveScanInterval))
	go func() {
		for range ticker.C {
			println("scanscanscanscan")
		}
	}()
}
