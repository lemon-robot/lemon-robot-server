package entity

import "time"

type DispatcherMachine struct {
	MachineCode       string    `gorm:"primary_key;size:64" json:"machine_code"`
	ArchTag           string    `gorm:"size:32" json:"arch_tag"`
	OsTag             string    `gorm:"size:32" json:"os_tag"`
	DispatcherVersion string    `gorm:"size:32" json:"dispatcher_version"`
	RegisterAt        time.Time `json:"register_at"`
	LastLoginAt       time.Time `json:"last_login_at"`
}
