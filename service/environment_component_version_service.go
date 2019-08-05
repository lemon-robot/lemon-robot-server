package service

import (
	"lemon-robot-server/dto"
)

type EnvironmentComponentVersionService interface {
	Save(been *dto.EnvironmentComponentVersionReq) (error, dto.EnvironmentComponentVersionReq)
	Delete(ecVersionKey string) error
	QueryList(environmentComponentKey string) (error, []dto.EnvironmentComponentVersionReq)
}
