package entity

import "time"

type Namespace struct {
	NamespaceKey       string `gorm:"primary_key;size:64" json:"namespace_key"`
	NamespaceTag       string `gorm:"index;size:64" json:"namespace_tag"`
	NamespaceName      string `gorm:"size:64" json:"namespace_name"`
	NamespaceIntroduce string `gorm:"size:1024" json:"namespace_introduce"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          *time.Time `sql:"index"`
}
