package entity

import "time"

type Task struct {
	LrEntity
	TaskKey            string    `gorm:"primary_key;size:64" json:"task_key"`
	BelongNamespace    Namespace `gorm:"ForeignKey:BelongNamespaceKey;AssociationForeignKey:NamespaceKey" json:"belong_namespace"`
	BelongNamespaceKey string    `json:"belong_namespace_key"`
	TaskName           string    `gorm:"size:64" json:"task_name"`
	TaskIntroduce      string    `gorm:"size:2048" json:"task_introduce"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          *time.Time `sql:"index"`
}
