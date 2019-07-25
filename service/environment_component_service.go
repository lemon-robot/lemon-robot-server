package service

import "lemon-robot-server/entity"

type EnvironmentService interface {
	Save(been *entity.EnvironmentComponent) (error, entity.EnvironmentComponent)
	Delete(key string) error
	QueryList() (error, []entity.EnvironmentComponent)
}
