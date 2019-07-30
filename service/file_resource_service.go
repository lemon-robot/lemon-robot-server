package service

import "mime/multipart"

type FileResourceService interface {
	Upload(multipart.File, *multipart.FileHeader) string
}
