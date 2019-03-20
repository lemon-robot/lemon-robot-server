package entity

type ParamValueDefine struct {
	ParamValueDefineKey string `gorm:"primary_key;size:64" json:"param_value_define_key"`
	BelongTaskKey       string `gorm:"index;size:64" json:"belong_task_key"`
	ParamValueTag       string `gorm:"size:64" json:"param_value_tag"`
	ParamValueName      string `gorm:"size:64" json:"param_value_name"`
	ParamValueIntroduce string `gorm:"size:512" json:"param_value_introduce"`
	ParamValueRegex     string `gorm:"size:512" json:"param_value_regex"`
}
