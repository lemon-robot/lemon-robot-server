package entity

import "github.com/jinzhu/gorm"

type Lrc struct {
	gorm.Model
	LrcKey     string `gorm:"index;size:64" json:"lrc_key"`
	Lrct       string `gorm:"index;size:64" json:"lrct"`
	LrcpSecret string `gorm:"size:512" json:"lrcp_secret"`
}
