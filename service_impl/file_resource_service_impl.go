package service_impl

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"lemon-robot-golang-commons/utils/lru_file"
	"lemon-robot-golang-commons/utils/lru_string"
	"lemon-robot-server/dao"
	"lemon-robot-server/entity"
	"lemon-robot-server/sysinfo"
	"mime/multipart"
	"os"
	"strings"
)

type FileResourceServiceImpl struct {
	fileResourceServiceDao *dao.FileResourceServiceDao
}

func NewFileResourceServiceImpl() *FileResourceServiceImpl {
	return &FileResourceServiceImpl{
		fileResourceServiceDao: dao.NewFileResourceServiceDao(),
	}
}

func (i *FileResourceServiceImpl) Upload(ctx *gin.Context, file multipart.File, handler *multipart.FileHeader) (string, error) {
	originalFileName, filePath, fileInfo, error := lru_file.GetInstance().SaveFileToTemporary(file, handler)
	if error != nil {
		return "", error
	}
	//log.Println("originalFileName == " + originalFileName)
	//log.Println("filePath == " + filePath)
	//log.Println("fileInfo.Size == " + strconv.FormatInt(fileInfo.Size(),10))

	uuidStr := lru_string.GetInstance().Uuid(true)
	error, fileSuffixName := lru_string.GetInstance().GetFileSuffixName(originalFileName)
	if error != nil {
		return "", error
	}
	fileName := uuidStr + fileSuffixName
	_, err := uploadFileToOSS(fileName, filePath)
	if err != nil {
		return "", err
	}
	been := entity.FileResource{}
	been.FileResourceKey = lru_string.GetInstance().Uuid(true)
	been.SourceType = fileSuffixName
	been.OriginalFileName = originalFileName
	been.FileExtension = fileSuffixName[1 :len(fileSuffixName)]
	been.FileSize = fileInfo.Size()
	been.FilePath = fileName
	contentTypeAll := ctx.Request.Header.Get("Content-Type")
	been.ContentType = contentTypeAll[: strings.Index(contentTypeAll, ";")]
	error = i.fileResourceServiceDao.Create(&been)
	if error != nil {
		return "", error
	}
	return been.FileResourceKey, nil
}

/**
 * 上传文件到oss中
 */
func uploadFileToOSS(fileName string, filePath string) (string, error) {
	config := sysinfo.LrServerConfig()
	s3 := config.FileResourceConfig
	creds := credentials.NewStaticCredentials(s3["secretId"], s3["secretKey"], "", )
	_, err := creds.Get()
	if err != nil {
		return "", errors.New("Credential error : " + err.Error())
	}
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials : creds,
		// 官网上面的地址是亚马逊的服务器, 这里区域改成自己的
		Region : aws.String(s3["secretKey"]),
		// 其他云的服务器
		Endpoint : aws.String(s3["requestFront"] + s3["endpoint"]),
		// 打印错误信息
		CredentialsChainVerboseErrors : aws.Bool(true),
	}))
	uploader := s3manager.NewUploader(sess)
	f, err  := os.Open(filePath)
	if err != nil {
		return "", errors.New("open file error : " + err.Error())
	}
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s3["bucket"]),
		Key:    aws.String(s3["rootPath"] + fileName),
		Body:   f,
	})
	if err != nil {
		return "", errors.New("uploadFile error : " + err.Error())
	}

	return "loacation : " + result.Location, nil
}



