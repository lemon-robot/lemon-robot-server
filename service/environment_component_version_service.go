package service

import "lemon-robot-server/entity"

type EnvironmentComponentVersionService interface {
	Save(been *entity.EnvironmentComponentVersion) (error, entity.EnvironmentComponentVersion)
	Delete(belongEnvironmentComponentKey string) error
	QueryList(environmentComponentKey string) (error, []entity.EnvironmentComponentVersion)
}
