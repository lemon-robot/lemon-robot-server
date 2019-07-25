package entity

import "time"

type EnvironmentComponent struct {
	EnvironmentComponentKey         string `gorm:"primary_key;size:64" json:"environmentComponentKey"`
	EnvironmentComponentName        string `gorm:"size:128" json:"environmentComponentName"`
	EnvironmentComponentDescription string `gorm:"size:10240" json:"environmentComponentDescription"`
	CreatedAt                       time.Time
	UpdatedAt                       time.Time
	DeletedAt                       *time.Time `sql:"index"`
}
