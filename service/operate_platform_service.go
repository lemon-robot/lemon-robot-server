package service

import "lemon-robot-server/dto"

type OperatePlatformService interface {
	GetAll() ([]dto.OperatePlatformReq, error)
	//GetOnes(key string) (dto.OperatePlatformReq, error)
}
