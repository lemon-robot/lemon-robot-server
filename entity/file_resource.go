package entity

import "time"

type FileResource struct {
	FileResourceKey  string `gorm:"primary_key;size:64" json:"fileResourceKey"`
	SourceType       string `gorm:"size:128" json:"sourceType"`
	OriginalFileName string `gorm:"size:512" json:"originalFileName"`
	FileExtension    string `gorm:"size:512" json:"fileExtension"`
	FileSize         int64  `json:"fileSize"`
	FilePath         string `gorm:"size:4096" json:"filePath"`
	ContentType      string `gorm:"size:128" json:"contentType"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeleteAt         *time.Time
}
