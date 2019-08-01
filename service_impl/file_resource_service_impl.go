package service_impl

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"lemon-robot-golang-commons/utils/lru_file"
	"lemon-robot-golang-commons/utils/lru_string"
	"lemon-robot-server/dao"
	"mime/multipart"
	"os"
)

//var baseFilePath = "/home/zgy/Applications/golang/file_resource_oss/"
var baseFilePath = os.TempDir()
var secretId = "AKIDQeeM8gG4wnNk48qfBm8F5fYHj1S1wZlt"
var secretKey = "Zsc6RYTEoDR4UjnkfMp20hFXqOkwsa5y"
var region = "cos.ap-beijing"
var endpoint = "https://cos.ap-beijing.myqcloud.com"
var bucket = "lemon-robot-server-1258459529"
var cosParentFile = "file-resource/"

type FileResourceServiceImpl struct {
	fileResourceServiceDao *dao.FileResourceServiceDao
}

func NewFileResourceServiceImpl() *FileResourceServiceImpl {
	return &FileResourceServiceImpl{
		fileResourceServiceDao: dao.NewFileResourceServiceDao(),
	}
}

func (i *FileResourceServiceImpl) Upload(file multipart.File, handler *multipart.FileHeader) (string, error) {
	error, fileName, filePath := lru_file.GetInstance().SaveFileToTemporary(file, handler)
	if error != nil {
		return "", error
	}
	uuidStr := lru_string.GetInstance().Uuid(true)
	error, fileSuffixName := lru_string.GetInstance().GetFileSuffixName(fileName)
	if error != nil {
		return "", error
	}
	fileName = uuidStr + fileSuffixName
	_, err := uploadOSS(fileName, filePath)
	if err != nil {
		return "", err
	}else {
		os.Remove(filePath)
		return fileName, nil
	}
}

/**
 * 上传文件到oss中
 */
func uploadOSS(fileName string, filePath string) (string, error) {
	creds := credentials.NewStaticCredentials(secretId, secretKey, "", )
	_, err := creds.Get()
	if err != nil {
		return "", errors.New("Credential error : " + err.Error())
	}
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials : creds,
		// 官网上面的地址是亚马逊的服务器, 这里区域改成自己的
		Region : aws.String(region),
		// 其他云的服务器
		Endpoint : aws.String(endpoint),
		// 打印错误信息
		CredentialsChainVerboseErrors : aws.Bool(true),
	}))
	uploader := s3manager.NewUploader(sess)
	f, err  := os.Open(filePath)
	if err != nil {
		return "", errors.New("open file error : " + err.Error())
	}
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(cosParentFile + fileName),
		Body:   f,
	})
	if err != nil {
		return "", errors.New("uploadFile error : " + err.Error())
	}

	return "loacation : " + result.Location, nil

}



