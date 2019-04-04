package entity

import "time"

type DispatcherTag struct {
	TagKey    string    `gorm:"primary_key;size:64" json:"tag_key"`
	TagName   string    `gorm:"size:32" json:"tag_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
