package service_server_node

import (
	"lemon-robot-golang-commons/logger"
	"lemon-robot-server/core/core_other"
	"lemon-robot-server/dao/dao_dispatcher_online"
	"lemon-robot-server/dao/dao_server_node"
	"lemon-robot-server/entity"
	"lemon-robot-server/sysinfo"
	"os"
	"runtime"
	"time"
)

var calculatedMachineCode string

func GetCalculatedMachineCode() string {
	if calculatedMachineCode == "" {
		machineCode, mcErr := core_other.GetMachineSign()
		if mcErr != nil {
			logger.Error("Server nodes cannot be registered because machine code cannot be computed", mcErr)
			os.Exit(1)
		}
		calculatedMachineCode = machineCode
	}
	return calculatedMachineCode
}

func RegisterServerNode() {
	dao_dispatcher_online.DeleteByClusterNodeMachineCode(GetCalculatedMachineCode())
	now := time.Now()
	nodeData := &entity.ServerNode{
		MachineCode:   GetCalculatedMachineCode(),
		CpuArch:       runtime.GOARCH,
		OperateSystem: runtime.GOOS,
		ServerVersion: sysinfo.AppVersion(),
		StartAt:       now,
		ActiveTime:    now,
	}
	dao_server_node.Save(nodeData)
}

func RefreshActiveTime() {
	dao_server_node.UpdateActiveTime(GetCalculatedMachineCode(), time.Now())
}
