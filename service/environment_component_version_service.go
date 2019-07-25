package service

import "lemon-robot-server/entity"

type EnvironmentComponentVersionService interface {
	Save(been *entity.EnvironmentComponentVersion) (error, entity.EnvironmentComponentVersion)
	Delete(belongEnvironmentComponentKey string) error
	QueryList() (error, []entity.EnvironmentComponentVersion)
}
