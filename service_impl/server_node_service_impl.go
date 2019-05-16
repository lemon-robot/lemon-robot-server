package service_impl

import (
	"fmt"
	"lemon-robot-golang-commons/utils/lru_date"
	"lemon-robot-golang-commons/utils/lru_machine"
	"lemon-robot-server/dao"
	"lemon-robot-server/dto"
	"lemon-robot-server/entity"
	"lemon-robot-server/sysinfo"
	"runtime"
	"time"
)

type ServerNodeServiceImpl struct {
	serverNodeDao         *dao.ServerNodeDao
	dispatcherOnlineDao   *dao.DispatcherOnlineDao
	calculatedMachineSign string
}

func NewServerNodeServiceImpl() *ServerNodeServiceImpl {
	return &ServerNodeServiceImpl{
		serverNodeDao:       dao.NewServerNodeDao(),
		dispatcherOnlineDao: dao.NewDispatcherOnlineDao(),
	}
}

func (i *ServerNodeServiceImpl) RegisterServerNode() {
	i.dispatcherOnlineDao.DeleteByServerNodeMachineSign(lru_machine.GetInstance().GetMachineSign())
	nodeData := i.serverNodeDao.FirstByExample(&entity.ServerNode{
		MachineSign: lru_machine.GetInstance().GetMachineSign(),
	})
	now := time.Now()
	if nodeData.MachineSign == "" {
		nodeData = entity.ServerNode{
			MachineSign:   lru_machine.GetInstance().GetMachineSign(),
			CpuArch:       runtime.GOARCH,
			OperateSystem: runtime.GOOS,
			ServerVersion: sysinfo.AppVersion(),
			StartAt:       now,
			ActiveTime:    now,
			Alias:         "",
		}
	} else {
		nodeData.StartAt = now
		nodeData.ActiveTime = now
	}
	i.serverNodeDao.Save(&nodeData)
}

func (i *ServerNodeServiceImpl) RefreshActiveTime() {
	i.serverNodeDao.UpdateActiveTime(lru_machine.GetInstance().GetMachineSign(), time.Now())
}

func (i *ServerNodeServiceImpl) UpdateAlias(machineSign, newAlias string) {
	i.serverNodeDao.UpdateAlias(machineSign, newAlias)
}

func (i *ServerNodeServiceImpl) QueryAllNodeInfo() []dto.ServerNodeResp {
	serverNodes := i.serverNodeDao.FindAllByExample(&entity.ServerNode{})
	serverNodeInfoArr := make([]dto.ServerNodeResp, len(serverNodes))
	for index, item := range serverNodes {
		serverNodeInfoArr[index] = dto.ServerNodeResp{
			NodeInfo: item,
			ActiveState: item.ActiveTime.After(lru_date.GetInstance().CalculateTimeByDurationStr(
				time.Now(), fmt.Sprintf("-%ds", sysinfo.LrServerConfig().ClusterNodeActiveInterval*2))),
		}
	}
	return serverNodeInfoArr
}
