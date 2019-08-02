package service_impl

import (
	"lemon-robot-server/dao"
	"lemon-robot-server/dto"
)

type OperatePlatformServiceImpl struct {
	operatePlatformDao *dao.OperatePlatformDao
}

func NewOperatePlatformServiceImpl() *OperatePlatformServiceImpl {
	return &OperatePlatformServiceImpl{
		operatePlatformDao : dao.NewOperatePlatformDao(),
	}
}

func (i *OperatePlatformServiceImpl) GetAll() ([]dto.OperatePlatformReq, error) {
	var operatePlatformReqs []dto.OperatePlatformReq
	operatePlatforms, error := i.operatePlatformDao.GetAll()
	if error != nil {
		return nil, error
	}
	for _, v := range operatePlatforms {
		operatePlatformReq := dto.OperatePlatformReq{}
		operatePlatformReq.OperatePlatformKey = v.OperatePlatformKey
		operatePlatformReq.OperatePlatformRemark = v.OperatePlatformRemark
		operatePlatformReq.CpuArchTag = v.CpuArchTag
		operatePlatformReq.OperateSystemTag = v.OperateSystemTag
		operatePlatformReqs = append(operatePlatformReqs, operatePlatformReq)
	}
	return operatePlatformReqs, nil
}

//func (i *OperatePlatformServiceImpl) GetOnes(key string) (dto.OperatePlatformReq, error) {
//	var operatePlatformReq dto.OperatePlatformReq
//	operatePlatform, error := i.operatePlatformDao.GetOnes(key)
//	if error != nil {
//		return operatePlatformReq, error
//	}
//	operatePlatformReq.OperatePlatformKey = operatePlatform.OperatePlatformKey
//	operatePlatformReq.OperateSystemTag = operatePlatform.OperateSystemTag
//	operatePlatformReq.CpuArchTag = operatePlatform.CpuArchTag
//	operatePlatformReq.OperatePlatformRemark = operatePlatform.OperatePlatformRemark
//	return operatePlatformReq, nil
//}