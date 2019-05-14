package entity

import "time"

type DispatcherTag struct {
	TagKey    string    `gorm:"primary_key;size:64" json:"tagKey"`
	TagName   string    `gorm:"size:32" json:"tagName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"-"`
}
