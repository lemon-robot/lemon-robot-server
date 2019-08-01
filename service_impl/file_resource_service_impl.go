package service_impl

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
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

func (i *FileResourceServiceImpl) Upload(file multipart.File, handler *multipart.FileHeader) (error, string) {
	error, fileName, filePath := saveFileDisk(file, handler)
	if error != nil {
		return error, ""
	}else {
		error, uploadResult := uploadOSS(fileName, filePath)
		if error != nil {
			return error, ""
		}else {
			os.Remove(filePath)
			return nil, uploadResult
		}
	}
}

/**
 * 保存文件到本地磁盘中
 */
func saveFileDisk(file multipart.File, handler *multipart.FileHeader) (error, string, string) {
	fileName := handler.Filename
	out, err := os.Create(baseFilePath + fileName)
	if err != nil {
		return err, "", ""
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return err, "", ""
	}
	return nil, fileName, baseFilePath + fileName
}

/**
 * 上传文件到oss中
 */
func uploadOSS(fileName string, filePath string) (error, string) {
	creds := credentials.NewStaticCredentials(secretId, secretKey, "", )
	_, err := creds.Get()
	if err != nil {
		return errors.New("Credential error : " + err.Error()), ""
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
		return errors.New("open file error : " + err.Error()), ""
	}
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(cosParentFile + fileName),
		Body:   f,
	})
	if err != nil {
		return errors.New("uploadFile error : " + err.Error()), ""
	}

	return nil, "loacation : " + result.Location

}



