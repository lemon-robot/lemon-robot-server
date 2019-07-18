package websocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/utils/lru_machine"
	"lemon-robot-golang-commons/utils/lru_string"
	"lemon-robot-server/entity"
	"lemon-robot-server/service"
	"lemon-robot-server/service_impl"
	"log"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var connectPool = make(map[string]*websocket.Conn)
var authService service.AuthService = service_impl.NewAuthServiceImpl()
var dispatcherMachineService = service_impl.NewDispatcherMachineServiceImpl()
var dispatcherOnlineService service.DispatcherOnlineService = service_impl.NewDispatcherOnlineServiceImpl()

func dealConnectionClose(onlineKey string) {
	delete(connectPool, onlineKey)
	dispatcherOnlineService.DeleteByOnlineKey(onlineKey)
	// TODO 删除正在此机器上执行的计划
	logger.Info(fmt.Sprint("A Websocket connection is broken, online key: ", onlineKey))
	logger.Info(fmt.Sprint("The number of remaining connections on the current node: ", len(connectPool)))
}

func ConnectHandler(context *gin.Context) {
	token := context.Param("token")
	onlineKey := lru_string.GetInstance().Uuid(true)
	if authService.CheckToken(token) {
		os := context.Param("os")
		arch := context.Param("arch")
		dispatcherVersion := context.Param("dispatcherVersion")
		machineSign := context.Param("machineSign")
		machineEntity := dispatcherMachineService.FindByServerNodeMachineSign(machineSign)
		if machineEntity.MachineSign == "" {
			machineEntity = entity.DispatcherMachine{
				MachineSign:       machineSign,
				CpuArch:           arch,
				OperateSystem:     os,
				DispatcherVersion: dispatcherVersion,
				Alias:             "",
				Tags:              make([]entity.DispatcherTag, 0),
			}
			dispatcherMachineService.Save(&machineEntity)
		}
		dispatcherOnlineService.Save(&entity.DispatcherOnline{
			OnlineKey:                 onlineKey,
			RelationMachineSign:       machineSign,
			RelationDispatcherMachine: machineEntity,
			IpAddress:                 context.ClientIP(),
			BindServerMachineSign:     lru_machine.GetInstance().GetMachineSign(),
		})
		logger.Info(
			fmt.Sprintf(
				"Websocket was successfully established! OS: %v, ARCH: %v, Dispatcher Ver: %v, Machine Sign: %v",
				os, arch, dispatcherVersion, machineSign))
	} else {
		logger.Warn("An illegal websocket connection request has been rejected by the system. token: " + token)
		return
	}
	c, err := upGrader.Upgrade(context.Writer, context.Request, nil)
	if c == nil || err != nil {
		logger.Error("Unable to create websocket connection properly", err)
		return
	}
	connectPool[onlineKey] = c
	logger.Info(fmt.Sprint("Websocket connection is successfully established! Online key: ", onlineKey))
	logger.Info(fmt.Sprint("The number of remaining connections on the current node: ", len(connectPool)))
	defer c.Close()
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			switch err.(type) {
			case *websocket.CloseError:
				dealConnectionClose(onlineKey)
			default:
				logger.Error("Errors occurred while reading cancelled messages from websocket", err)
			}
			break
		}
		log.Printf("Receive messages from websocket: %s", message)
	}
}
