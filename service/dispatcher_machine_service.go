package service

import "lemon-robot-server/entity"

type DispatcherMachineService interface {
	Save(dispatcherMachine *entity.DispatcherMachine)
}
