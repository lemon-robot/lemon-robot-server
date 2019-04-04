package service_server_node

import (
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/utils/machine"
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
	return calculatedMachineCode
}

func RegisterServerNode() {
	machineCode, mcErr := lrumachine.CalculateMachineCode()
	if mcErr != nil {
		logger.Error("Server nodes cannot be registered because machine code cannot be computed", mcErr)
		os.Exit(1)
	}
	//DELETE FROM `lr_dispatcher_onlines`  WHERE (bind_server_machine_code = '537cc9f908067d8066b9129b7cc20499')
	dao_dispatcher_online.Delete("bind_server_machine_code = ?", machineCode)
	//db.Db().Where("bind_server_machine_code = ?", machineCode).Delete(&entity.DispatcherOnline{})
	calculatedMachineCode = machineCode
	now := time.Now()
	dao_server_node.Save(&entity.ServerNode{
		MachineCode:   machineCode,
		CpuArch:       runtime.GOARCH,
		OperateSystem: runtime.GOOS,
		ServerVersion: sysinfo.AppVersion(),
		StartAt:       now,
		ActiveTime:    now,
	})
}

func ClearNotActiveServerNodes() {

}

func RefreshActiveTime() {

}
