package entity

import "time"

type OperatePlatform struct {
	OperatePlatformKey    string `gorm:"size:64;unique_index" json:"operatePlatformKey"`
	OperateSystemTag      string `gorm:"size:32;primary_key" json:"operateSystemTag"`
	CpuArchTag            string `gorm:"size:32;primary_key" json:"cpuArchTag"`
	OperatePlatformRemark string `gorm:"size:2048" json:"operatePlatformRemark"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             *time.Time `sql:"index"`
}
