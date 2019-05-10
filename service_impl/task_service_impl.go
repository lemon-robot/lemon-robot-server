package service_impl

import (
	"lemon-robot-server/core/gitter"
	"lemon-robot-server/entity"
)

type TaskServiceImpl struct{}

func NewTaskServiceImpl() *TaskServiceImpl {
	return &TaskServiceImpl{}
}

func (i *TaskServiceImpl) Create(task entity.Task) error {
	err := gitter.Ins().TaskCreate(&task)
	return err
}
