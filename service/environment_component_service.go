package service

import (
	"lemon-robot-server/dto"
)

type EnvironmentService interface {
	Save(been *dto.EnvironmentComponentReq) (error, dto.EnvironmentComponentReq)
	Delete(key string) error
	QueryList() (error, []dto.EnvironmentComponentReq)
}
