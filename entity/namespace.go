package entity

import "time"

type Namespace struct {
	NamespaceKey       string `gorm:"primary_key;size:64" json:"namespaceKey"`
	NamespaceTag       string `gorm:"index;size:64" json:"namespaceTag"`
	NamespaceName      string `gorm:"size:64" json:"namespaceName"`
	NamespaceIntroduce string `gorm:"size:1024" json:"nameSpaceIntroduce"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          *time.Time `sql:"index"`
}
