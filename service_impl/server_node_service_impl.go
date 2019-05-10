package service_impl

import (
	"fmt"
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
	now := time.Now()
	nodeData := &entity.ServerNode{
		MachineSign:   lru_machine.GetInstance().GetMachineSign(),
		CpuArch:       runtime.GOARCH,
		OperateSystem: runtime.GOOS,
		ServerVersion: sysinfo.AppVersion(),
		StartAt:       now,
		ActiveTime:    now,
	}
	i.serverNodeDao.Save(nodeData)
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
	dur, _ := time.ParseDuration(fmt.Sprintf("-%ds", sysinfo.LrServerConfig().ClusterNodeActiveInterval*2))
	for index, item := range serverNodes {
		serverNodeInfoArr[index] = dto.ServerNodeResp{
			NodeInfo:    item,
			ActiveState: item.ActiveTime.After(time.Now().Add(dur)),
		}
	}
	return serverNodeInfoArr
}
