package entity

import "time"

type FileResource struct {
	ID               uint `gorm:"primary_key"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	FileResourceKey  string `gorm:"index;size:64" json:"file_resource_key"`
	UsedTag          bool   `json:"used_tag"`
	UploadedTag      bool   `json:"uploaded_tag"`
	OriginalFileName string `gorm:"size:512" json:"original_file_name"`
	FileExtension    string `gorm:"size:512" json:"file_extension"`
	FileSize         int64  `json:"file_size"`
	ContentType      string `json:"content_type"`
}
