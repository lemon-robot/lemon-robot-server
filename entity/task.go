package entity

import "time"

type Task struct {
	TaskKey              string             `gorm:"primary_key;size:64" json:"task_key"`
	BelongNamespace      Namespace          `gorm:"ForeignKey:BelongNamespaceKey;AssociationForeignKey:NamespaceKey" json:"belong_namespace"`
	BelongNamespaceKey   string             `gorm:"size:64;unique_index:task_unique_index" json:"belong_namespace_key"`
	TaskTag              string             `gorm:"size:64;unique_index:task_unique_index" json:"task_tag"`
	TaskName             string             `gorm:"size:64" json:"task_name"`
	TaskIntroduce        string             `gorm:"size:2048" json:"task_introduce"`
	ParamValueDefineList []ParamValueDefine `gorm:"ForeignKey:BelongTaskKey" json:"param_value_define_list"`
	ParamFileDefineList  []ParamFileDefine  `gorm:"ForeignKey:BelongTaskKey" json:"param_file_define_list"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            *time.Time `sql:"index"`
}
