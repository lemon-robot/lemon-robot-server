package service

import "lemon-robot-server/dto"

type ServerNodeService interface {
	RegisterServerNode()
	RefreshActiveTime()
	UpdateAlias(machineSign, newAlias string)
	QueryAllNodeInfo() []dto.ServerNodeResp
}
