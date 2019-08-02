package service

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type FileResourceService interface {
	Upload(*gin.Context, multipart.File, *multipart.FileHeader) (string, error)
}
