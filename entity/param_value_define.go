package entity

type ParamValueDefine struct {
	ParamValueDefineKey string `gorm:"primary_key;size:64" json:"paramValueDefineKey"`
	BelongTaskKey       string `gorm:"index;size:64" json:"belongTaskKey"`
	ParamValueTag       string `gorm:"size:64" json:"paramValueTag"`
	ParamValueName      string `gorm:"size:64" json:"paramValueName"`
	ParamValueIntroduce string `gorm:"size:512" json:"paramValueIntroduce"`
	ParamValueRegex     string `gorm:"size:512" json:"paramValueRegex"`
}
