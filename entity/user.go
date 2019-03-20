package entity

import "time"

type User struct {
	UserKey        string `gorm:"primary_key;size:64" json:"user_key"`
	UserNumber     string `gorm:"unique_index;size:64" json:"user_number"`
	PasswordSecret string `gorm:"size:512" json:"password_secret"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `sql:"index"`
}
