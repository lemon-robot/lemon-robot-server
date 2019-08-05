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

type EnvironmentComponentVersionServiceImpl struct {
	environmentComponentVersionDao *dao.EnvironmentComponentVersionDao
}

func NewEnvironmentComponentVersionServiceImpl() *EnvironmentComponentVersionServiceImpl {
	return &EnvironmentComponentVersionServiceImpl{
		environmentComponentVersionDao : dao.NewEnvironmentComponentVersionDao(),
	}
}

func (i *EnvironmentComponentVersionServiceImpl) Save(been *dto.EnvironmentComponentVersionReq) (error, dto.EnvironmentComponentVersionReq) {
	ecVersionKey := been.ECVersionKey
	environmentComponentVersion := entity.EnvironmentComponentVersion{}
	environmentComponentVersion.BelongEnvironmentComponentKey = been.BelongEnvironmentComponentKey
	environmentComponentVersion.BindOperatePlatformKey = been.BindOperatePlatformKey
	environmentComponentVersion.ECVersionDescription = been.ECVersionDescription
	environmentComponentVersion.ECVersionName = been.ECVersionName
	environmentComponentVersion.ECVersionNumber = been.ECVersionNumber
	environmentComponentVersion.ECVersionTag = been.ECVersionTag
	environmentComponentVersion.StateCheckScript = been.StateCheckScript
	environmentComponentVersion.InstallScript = been.InstallScript
	environmentComponentVersion.UninstallScript = been.UninstallScript
	environmentComponentVersion.ProgramFileResourceKey = been.ProgramFileResourceKey
	environmentComponentVersion.WhereToInstall = been.WhereToInstall
	environmentComponentVersion.BindOperatePlatformKey = been.BindOperatePlatformKey
	environmentComponentVersion.ECVersionKey = ecVersionKey
	if ecVersionKey == "" {
		been.ECVersionKey = lru_string.GetInstance().Uuid(true)
		environmentComponentVersion.ECVersionKey = been.ECVersionKey
		error, _ :=i.environmentComponentVersionDao.Create(&environmentComponentVersion)
		return error, *been
	}else {
		error, _ := i.environmentComponentVersionDao.QueryOne(environmentComponentVersion.ECVersionKey)
		if error != nil {
			return error, *been
		}
		err, _ := i.environmentComponentVersionDao.Updata(&environmentComponentVersion)
		return err, *been
	}
}

func (i *EnvironmentComponentVersionServiceImpl) Delete(ecVersionKey string) error {
	error, queryOneiEnvironmentComponentVersion := i.environmentComponentVersionDao.QueryOne(ecVersionKey)
	if error != nil {
		return error
	}
	return i.environmentComponentVersionDao.Delete(&queryOneiEnvironmentComponentVersion, time.Now())
}

func (i *EnvironmentComponentVersionServiceImpl) QueryList(belongEnvironmentComponentKey string) (error, []dto.EnvironmentComponentVersionReq) {
	error, environmentComponentVersions := i.environmentComponentVersionDao.QueryList(belongEnvironmentComponentKey)
	var environmentComponentVersionReqs []dto.EnvironmentComponentVersionReq
	fileResourceConfig := sysinfo.LrServerConfig().FileResourceConfig
	for _, v := range environmentComponentVersions {
		environmentComponentVersion := dto.EnvironmentComponentVersionReq{}
		environmentComponentVersion.ECVersionKey = v.ECVersionKey
		environmentComponentVersion.BelongEnvironmentComponentKey = belongEnvironmentComponentKey
		// 嵌套对象:environmentComponent
		_, environmentComponent := dao.NewEnvironmentComponentDao().QueryOne(belongEnvironmentComponentKey)
		environmentComponentReq := dto.EnvironmentComponentReq{}
		environmentComponentReq.EnvironmentComponentKey = environmentComponent.EnvironmentComponentKey
		environmentComponentReq.IconFileResourceKey = environmentComponent.IconFileResourceKey
		environmentComponentReq.EnvironmentComponentName = environmentComponent.EnvironmentComponentName
		environmentComponentReq.EnvironmentComponentDescription = environmentComponent.EnvironmentComponentDescription
		environmentComponentReq.EnvironmentComponentVersionCount = len(environmentComponentVersions)
		// 取iconUrl
		fileResourceDao := dao.NewFileResourceServiceDao()
		environmentComponentFileSource, _:= fileResourceDao.GetFileSource(environmentComponent.IconFileResourceKey)
		environmentComponentIconPath := environmentComponentFileSource.FilePath
		environmentComponentReq.IconFileResourceUrl = fileResourceConfig["requestFront"] + fileResourceConfig["bucket"] + "." + fileResourceConfig["endpoint"] + string(os.PathSeparator) + fileResourceConfig["rootPath"] + environmentComponentIconPath
		environmentComponentVersion.BelongEnvironmentComponent = environmentComponentReq

		environmentComponentVersion.BindOperatePlatformKey = v.BindOperatePlatformKey
		// 嵌套对象:OperatePlatform
		operatePlatform,_ := dao.NewOperatePlatformDao().GetOnes(v.BindOperatePlatformKey)
		operatePlatformReq := dto.OperatePlatformReq{}
		operatePlatformReq.OperatePlatformRemark = operatePlatform.OperatePlatformRemark
		operatePlatformReq.CpuArchTag = operatePlatform.CpuArchTag
		operatePlatformReq.OperateSystemTag = operatePlatform.OperateSystemTag
		operatePlatformReq.OperatePlatformKey = operatePlatform.OperatePlatformKey
		environmentComponentVersion.BindOperatePlatform = operatePlatformReq

		environmentComponentVersion.ECVersionDescription = v.ECVersionDescription
		environmentComponentVersion.ECVersionName = v.ECVersionName
		environmentComponentVersion.ECVersionNumber = v.ECVersionNumber
		environmentComponentVersion.ECVersionTag = v.ECVersionTag
		environmentComponentVersion.StateCheckScript = v.StateCheckScript
		environmentComponentVersion.InstallScript = v.InstallScript
		environmentComponentVersion.UninstallScript = v.UninstallScript
		// 嵌套对象:FileResource
		environmentComponentVersion.ProgramFileResourceKey = v.ProgramFileResourceKey
		fileResource,_ := fileResourceDao.GetFileSource(environmentComponentVersion.ProgramFileResourceKey)
		fileResourceReq := dto.FileResourceReq{}
		fileResourceReq.FilePath = fileResource.FilePath
		fileResourceReq.FileResourceKey = fileResource.FileResourceKey
		fileResourceReq.ContentType = fileResource.ContentType
		fileResourceReq.OriginalFileName = fileResource.OriginalFileName
		fileResourceReq.SourceType = fileResource.SourceType
		fileResourceReq.FileExtension = fileResource.FileExtension
		fileResourceReq.FileSize = fileResource.FileSize
		environmentComponentVersion.ProgramFileResource = fileResourceReq

		environmentComponentVersion.WhereToInstall = v.WhereToInstall

		environmentComponentVersionReqs = append(environmentComponentVersionReqs, environmentComponentVersion)
	}
	return error, environmentComponentVersionReqs
}