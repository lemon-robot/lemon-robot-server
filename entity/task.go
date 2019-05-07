package entity

import "time"

type Task struct {
	TaskKey              string             `gorm:"primary_key;size:64" json:"taskKey"`
	BelongNamespace      Namespace          `gorm:"ForeignKey:BelongNamespaceKey;AssociationForeignKey:NamespaceKey" json:"belongNamespace"`
	BelongNamespaceKey   string             `gorm:"size:64;unique_index:task_unique_index" json:"belongNamespaceKey"`
	TaskTag              string             `gorm:"size:64;unique_index:task_unique_index" json:"taskTag"`
	TaskName             string             `gorm:"size:64" json:"taskName"`
	TaskIntroduce        string             `gorm:"size:2048" json:"taskIntroduce"`
	ParamValueDefineList []ParamValueDefine `gorm:"ForeignKey:BelongTaskKey" json:"paramValueDefineList"`
	ParamFileDefineList  []ParamFileDefine  `gorm:"ForeignKey:BelongTaskKey" json:"paramFileDefineList"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            *time.Time `sql:"index"`
}
