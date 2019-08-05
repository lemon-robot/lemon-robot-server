package service_impl

import (
	"lemon-robot-golang-commons/utils/lru_string"
	"lemon-robot-server/dao"
	"lemon-robot-server/dto"
	"lemon-robot-server/entity"
	"lemon-robot-server/sysinfo"
	"os"
	"time"
)

type EnvironmentComponentServiceImpl struct {
	environmentComponentDao *dao.EnvironmentComponentDao
}

func NewEnvironmentComponentServiceImpl() *EnvironmentComponentServiceImpl {
	return &EnvironmentComponentServiceImpl{
		environmentComponentDao : dao.NewEnvironmentComponentDao(),
	}
}

func (i *EnvironmentComponentServiceImpl) Save(environmentComponentReq *dto.EnvironmentComponentReq) (error, dto.EnvironmentComponentReq) {
	environmentComponent := entity.EnvironmentComponent{}
	environmentComponent.EnvironmentComponentKey = environmentComponentReq.EnvironmentComponentKey
	environmentComponent.EnvironmentComponentDescription = environmentComponentReq.EnvironmentComponentDescription
	environmentComponent.EnvironmentComponentName = environmentComponentReq.EnvironmentComponentName
	environmentComponent.IconFileResourceKey = environmentComponentReq.IconFileResourceKey
	if environmentComponent.EnvironmentComponentKey == "" {
		environmentComponent.EnvironmentComponentKey = lru_string.GetInstance().Uuid(true)
		environmentComponentReq.EnvironmentComponentKey = environmentComponent.EnvironmentComponentKey
		error, _ := i.environmentComponentDao.Create(&environmentComponent)
		return error, *environmentComponentReq
	}else {
		error, _ := i.environmentComponentDao.QueryOne(environmentComponentReq.EnvironmentComponentKey)
		if error != nil {
			return error, *environmentComponentReq
		}
		err, _ := i.environmentComponentDao.Update(&environmentComponent)
		return err, *environmentComponentReq
	}
}

func (i *EnvironmentComponentServiceImpl) Delete(key string) error {
	error, queryEnviromentComponent := i.environmentComponentDao.QueryOne(key)
	if error != nil {
		return error
	}
	currentTime:=time.Now()
	return i.environmentComponentDao.Delete(&queryEnviromentComponent, currentTime)
}

func (i *EnvironmentComponentServiceImpl) QueryList() (error, []dto.EnvironmentComponentReq) {
	var environmentComponentReqs []dto.EnvironmentComponentReq
	error, environmentComponents := i.environmentComponentDao.QueryList()
	fileResourceConfig := sysinfo.LrServerConfig().FileResourceConfig
	for _, v := range environmentComponents {
		environmentComponentReq := dto.EnvironmentComponentReq{}
		environmentComponentReq.EnvironmentComponentKey = v.EnvironmentComponentKey
		environmentComponentReq.EnvironmentComponentName = v.EnvironmentComponentName
		environmentComponentReq.EnvironmentComponentDescription = v.EnvironmentComponentDescription
		environmentComponentReq.IconFileResourceKey = v.IconFileResourceKey
		environmentComponentReq.EnvironmentComponentVersionCount = i.environmentComponentDao.QueryVersionCount(v.EnvironmentComponentKey)
		// 拼接IconFileResourceUrl, 返回的是能直接下载的地址
		fileSource, _:= dao.NewFileResourceServiceDao().GetFileSource(v.IconFileResourceKey)
		filePath := fileSource.FilePath
		environmentComponentReq.IconFileResourceUrl = fileResourceConfig["requestFront"] + fileResourceConfig["bucket"] + "." + fileResourceConfig["endpoint"] + string(os.PathSeparator) + fileResourceConfig["rootPath"] + filePath
		environmentComponentReqs = append(environmentComponentReqs, environmentComponentReq)
	}
	return error, environmentComponentReqs
}