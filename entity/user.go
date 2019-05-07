package entity

import "time"

type User struct {
	UserKey        string `gorm:"primary_key;size:64" json:"userKey"`
	UserNumber     string `gorm:"unique_index;size:64" json:"userNumber"`
	PasswordSecret string `gorm:"size:512" json:"passwordSecret"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `sql:"index"`
}
