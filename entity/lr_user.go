package entity

import "github.com/jinzhu/gorm"

type LrUser struct {
	gorm.Model
	LrUserKey      string `gorm:"index;size:64" json:"lr_user_key"`
	UserNumber     string `gorm:"index;size:64" json:"user_number"`
	PasswordSecret string `gorm:"size:512" json:"password_secret"`
}
