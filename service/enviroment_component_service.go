package service

import "lemon-robot-server/entity"

type EnviromentService interface {
	Save(been *entity.EnvironmentComponent) (error, entity.EnvironmentComponent)
	Delete(key string) error
	Query(key string) (error, entity.EnvironmentComponent)
}
