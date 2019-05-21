package service

import (
	"lemon-robot-server/dto"
	"lemon-robot-server/entity"
)

type DispatcherMachineService interface {
	SetAlias(req *dto.CommonMachineSetAliasReq) bool
	SetTags(req *dto.DispatcherMachineSetTagsReq) bool
	Save(dispatcherMachine *entity.DispatcherMachine)
}
