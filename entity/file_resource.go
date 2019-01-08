package entity

import "github.com/jinzhu/gorm"

type FileResource struct {
	gorm.Model
	FileKey          string `gorm:"index;size:64" json:"file_key"`
	UsedTag          bool   `json:"used_tag"`
	UploadedTag      bool   `json:"uploaded_tag"`
	OriginalFileName string `gorm:"size:512" json:"original_file_name"`
	FileExtension    string `gorm:"size:512" json:"file_extension"`
}
