package service

import (
	"lemon-robot-server/dto"
)

type DispatcherTagService interface {
	Save(tagSaveReq *dto.DispatcherTagSaveReq)
	Delete(tagKey string)
	List() []dto.DispatcherTagSaveReq
}
