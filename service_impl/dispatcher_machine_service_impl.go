package service_impl

import (
	"lemon-robot-server/dao"
	"lemon-robot-server/entity"
)

type DispatcherMachineServiceImpl struct {
	dispatcherMachineDao *dao.DispatcherMachineDao
}

func NewDispatcherMachineServiceImpl() *DispatcherMachineServiceImpl {
	return &DispatcherMachineServiceImpl{
		dispatcherMachineDao: dao.NewDispatcherMachineDao(),
	}
}

func (i *DispatcherMachineServiceImpl) Save(dispatcherMachine *entity.DispatcherMachine) {
	i.dispatcherMachineDao.Save(dispatcherMachine)
}
