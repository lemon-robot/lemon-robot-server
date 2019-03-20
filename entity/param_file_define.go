package entity

type ParamFileDefine struct {
	ParamFileDefineKey string `gorm:"primary_key;size:64" json:"param_file_define_key"`
	BelongTaskKey      string `gorm:"index;size:64" json:"belong_task_key"`
	ParamFileTag       string `gorm:"size:64" json:"param_file_tag"`
	ParamFileName      string `gorm:"size:64" json:"param_file_name"`
	ParamFileIntroduce string `gorm:"size:512" json:"param_file_introduce"`
}
