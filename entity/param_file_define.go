package entity

type ParamFileDefine struct {
	ParamFileDefineKey string `gorm:"primary_key;size:64" json:"paramFileDefineKey"`
	BelongTaskKey      string `gorm:"index;size:64" json:"belongTaskKey"`
	ParamFileTag       string `gorm:"size:64" json:"paramFileTag"`
	ParamFileName      string `gorm:"size:64" json:"paramFileName"`
	ParamFileIntroduce string `gorm:"size:512" json:"paramFileIntroduce"`
}
