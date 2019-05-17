package service

import (
	"lemon-robot-server/dto"
	"lemon-robot-server/entity"
)

type DispatcherTagService interface {
	Save(tagSaveReq *dto.DispatcherTagSaveReq)
	Delete(tagKey string)
	List() []entity.DispatcherTag
}
