package service_server_node

import (
	"fmt"
	"lemon-robot-golang-commons/logger"
	lrumachine "lemon-robot-golang-commons/utils/machine"
	"lemon-robot-server/dao/dao_dispatcher_online"
	"lemon-robot-server/dao/dao_server_node"
	"lemon-robot-server/dto"
	"lemon-robot-server/entity"
	"lemon-robot-server/sysinfo"
	"os"
	"runtime"
	"time"
)

type ServerNodeService struct {
	serverNodeDao         dao_dispatcher_online.DispatcherOnlineDao
	calculatedMachineSign string
}

func (i ServerNodeService) GetCalculatedMachineSign() string {
	if i.calculatedMachineSign == "" {
		machineSign, mcErr := lrumachine.GetMachineSign()
		if mcErr != nil {
			logger.Error("Server nodes cannot be registered because machine sign cannot be computed", mcErr)
			os.Exit(1)
		}
		i.calculatedMachineSign = machineSign
	}
	return i.calculatedMachineSign
}

func (i ServerNodeService) RegisterServerNode() {
	i.serverNodeDao.DeleteByClusterNodeMachineSign(i.GetCalculatedMachineSign())
	now := time.Now()
	nodeData := &entity.ServerNode{
		MachineSign:   i.GetCalculatedMachineSign(),
		CpuArch:       runtime.GOARCH,
		OperateSystem: runtime.GOOS,
		ServerVersion: sysinfo.AppVersion(),
		StartAt:       now,
		ActiveTime:    now,
	}
	dao_server_node.Save(nodeData)
}

func (i ServerNodeService) RefreshActiveTime() {
	dao_server_node.UpdateActiveTime(i.GetCalculatedMachineSign(), time.Now())
}

func (i ServerNodeService) UpdateAlias(machineSign, newAlias string) {
	dao_server_node.UpdateAlias(machineSign, newAlias)
}

func (i ServerNodeService) QueryAllNodeInfo() []dto.ServerNodeResp {
	serverNodes := dao_server_node.FindAllByExample(&entity.ServerNode{})
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
