package entity

import "time"

type FileResource struct {
	FileResourceKey  string `gorm:"primary_key;size:64" json:"file_resource_key"`
	UsedTag          bool   `json:"used_tag"`
	UploadedTag      bool   `json:"uploaded_tag"`
	OriginalFileName string `gorm:"size:512" json:"original_file_name"`
	FileExtension    string `gorm:"size:512" json:"file_extension"`
	FileSize         int64  `json:"file_size"`
	ContentType      string `gorm:"size:128" json:"content_type"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
