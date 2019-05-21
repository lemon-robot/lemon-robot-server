package service_impl

import (
	"lemon-robot-server/dao"
	"lemon-robot-server/db"
	"lemon-robot-server/dto"
	"lemon-robot-server/entity"
)

type DispatcherMachineServiceImpl struct {
	dispatcherMachineDao *dao.DispatcherMachineDao
	dispatcherTagDao     *dao.DispatcherTagDao
}

func NewDispatcherMachineServiceImpl() *DispatcherMachineServiceImpl {
	return &DispatcherMachineServiceImpl{
		dispatcherMachineDao: dao.NewDispatcherMachineDao(),
	}
}

func (i *DispatcherMachineServiceImpl) SetAlias(req *dto.CommonMachineSetAliasReq) bool {
	dispatcherMachine := i.dispatcherMachineDao.FirstByExample(&entity.DispatcherMachine{
		MachineSign: req.MachineSign,
	})
	if dispatcherMachine.MachineSign != "" {
		dispatcherMachine.Alias = req.Alias
		if i.dispatcherMachineDao.Save(&dispatcherMachine).Error == nil {
			return true
		}
	}
	return false
}

func (i *DispatcherMachineServiceImpl) SetTags(req *dto.DispatcherMachineSetTagsReq) bool {
	dispatcherMachineEntity := i.dispatcherMachineDao.FirstByExample(&entity.DispatcherMachine{
		MachineSign: req.MachineSign,
	})
	if dispatcherMachineEntity.MachineSign != "" {
		tags := make([]entity.DispatcherTag, 0)
		for _, tagKey := range req.TagKeys {
			tagEntity := i.dispatcherTagDao.FirstByExample(&entity.DispatcherTag{
				TagKey: tagKey,
			})
			if tagEntity.TagKey != "" {
				tags = append(tags, tagEntity)
			}
		}
		db.Db().Model(&dispatcherMachineEntity).Association("Tags").Replace(&tags)
		return true
	}
	return false
}

func (i *DispatcherMachineServiceImpl) Save(dispatcherMachine *entity.DispatcherMachine) {
	i.dispatcherMachineDao.Save(dispatcherMachine)
}

func (i *DispatcherMachineServiceImpl) FindByServerNodeMachineSign(machineSign string) entity.DispatcherMachine {
	return i.dispatcherMachineDao.FirstByExample(&entity.DispatcherMachine{
		MachineSign: machineSign,
	})
}
