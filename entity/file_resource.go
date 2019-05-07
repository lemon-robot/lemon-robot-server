package entity

import "time"

type FileResource struct {
	FileResourceKey  string `gorm:"primary_key;size:64" json:"fileResourceKey"`
	UsedTag          bool   `json:"usedTag"`
	UploadedTag      bool   `json:"uploadedTag"`
	OriginalFileName string `gorm:"size:512" json:"originalFileName"`
	FileExtension    string `gorm:"size:512" json:"fileExtension"`
	FileSize         int64  `json:"fileSize"`
	FileContent      []byte `gorm:"size:2048000" json:"fileContent"`
	ContentType      string `gorm:"size:128" json:"contentType"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
