package service_impl

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"fmt"
	"io"
	"io/ioutil"
	"lemon-robot-server/dao"
	"log"
	"mime/multipart"
	"os"
)

var baseFilePath = "/home/zgy/Applications/golang/file_resource_oss/"

type FileResourceServiceImpl struct {
	fileResourceServiceDao *dao.FileResourceServiceDao
}

func NewFileResourceServiceImpl() *FileResourceServiceImpl {
	return &FileResourceServiceImpl{
		fileResourceServiceDao: dao.NewFileResourceServiceDao(),
	}
}

func (i *FileResourceServiceImpl) Upload(file multipart.File, handler *multipart.FileHeader) string {
	save, fileName, filePath := saveFileDisk(file, handler)
	if save {
		result := uploadOSS(fileName, filePath)
		fmt.Println("result === ", result)
		return result
	} else {
		return "save file fail"
	}
}

/**
 * 保存文件到本地磁盘中
 */
func saveFileDisk(file multipart.File, handler *multipart.FileHeader) (saveResult bool, fileName string, filePath string) {
	filename := handler.Filename
	out, err := os.Create(baseFilePath + filename)
	if err != nil {
		log.Fatal(err)
		return false, "", ""
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
		return false, "", ""
	}
	return true, filename, baseFilePath + filename
}

/**
 * 上传文件到oss中
 */
func uploadOSS(fileName string, filePath string) (fileLoadUrl string) {

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(endpoints.ApSoutheast1RegionID),
	}))
	service := s3.New(sess)

	fp, err := os.Open(filePath)
	if err != nil {
		fmt.Println("err ==== ", err.Error())
	}
	//So(err, ShouldBeNil)
	defer fp.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()

	_, err = service.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String("lemon-robot-server-1258459529"),
		Key:    aws.String("file-resource/" + fileName),
		Body:   fp,
	})
	if err != nil {
		fmt.Println("upload oss success")
	}
	return "upload oss success"

	//So(err, ShouldBeNil)

	//byteArr, err := readAll(filePath)
	//if err != nil {
	//	return "Error reading file"
	//}
	//leanCloudUrl := "https://foqprysq.api.lncld.net/1.1/files/test.png"
	//req, err := http.NewRequest("POST", leanCloudUrl, bytes.NewBuffer(byteArr))
	//req.Header.Set("X-LC-Id", "FOqprYsqFNQ2vlQh5cobzROX-gzGzoHsz")
	//req.Header.Set("X-LC-Key", "iiIkDetwgoSvD988hNLW60pU")
	//req.Header.Set("Content-Type", "image/png")
	//client := &http.Client{}
	//resp, err := client.Do(req)
	//
	//defer resp.Body.Close()
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	// handle error
	//}
	//
	//fmt.Println("请求结果 ================ ", string(body))
	//return string(body)
}

/**
 * 获取s3服务
 */
//func getS3Service() *S3 {
//	sess := session.Must(session.NewSession(&aws.Config{
//		Region: aws.String(endpoints.ApSoutheast1RegionID),
//	}))
//	return s3.New(sess)
//}

func readAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}
