package service_task

import (
	"lemon-robot-server/core/gitter"
	"lemon-robot-server/entity"
)

func Create(task entity.Task) error {
	err := gitter.Ins().TaskCreate(&task)
	return err
}
