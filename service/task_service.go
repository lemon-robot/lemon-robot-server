package service

import "lemon-robot-server/entity"

type TaskService interface {
	Create(task entity.Task) error
}
