package entity

import "time"

type EnvironmentComponentDependencyRelation struct {
	ECVersionKey               string `gorm:"primary_key;size:64" json:"ecVersionKey"`
	DepEnvironmentComponentKey string `gorm:"size:64" json:"depEnvironmentComponentKey"`
	DepECVersionNumberRule     string `gorm:"size:128" json:"depEcVersionNumberRule"`
	CreatedAt                  time.Time
}
