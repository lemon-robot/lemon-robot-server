package service_impl

import (
	"lemon-robot-golang-commons/utils/lru_string"
	"lemon-robot-server/dao"
	"lemon-robot-server/dto"
	"lemon-robot-server/entity"
)

type DispatcherTagServiceImpl struct {
	dispatcherTagDao     *dao.DispatcherTagDao
	dispatcherMachineDao *dao.DispatcherMachineDao
}

func NewDispatcherTagServiceImpl() *DispatcherTagServiceImpl {
	return &DispatcherTagServiceImpl{
		dispatcherTagDao:     dao.NewDispatcherTagDao(),
		dispatcherMachineDao: dao.NewDispatcherMachineDao(),
	}
}

func (i *DispatcherTagServiceImpl) Save(tagSaveReq *dto.DispatcherTagSaveReq) {
	tagEntity := i.dispatcherTagDao.FirstByExample(&entity.DispatcherTag{
		TagKey: tagSaveReq.TagKey,
	})
	tagEntity.TagName = tagSaveReq.TagName
	if tagSaveReq.TagKey == "" || tagEntity.TagKey == "" {
		tagEntity.TagKey = lru_string.GetInstance().Uuid(true)
	}
	i.dispatcherTagDao.Save(&tagEntity)
}

func (i *DispatcherTagServiceImpl) Delete(tagKey string) {
	i.dispatcherTagDao.Delete("tag_key = ?", tagKey)
}

func (i *DispatcherTagServiceImpl) List() []entity.DispatcherTag {
	return i.dispatcherTagDao.FindAllByExample(&entity.DispatcherTag{})
}
