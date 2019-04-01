package entity

import "time"

type DispatcherInstance struct {
	InstanceKey string    `gorm:"primary_key;size:64" json:"instance_key"`
	IpAddress   string    `gorm:"size:64" json:"ip_address"`
	BindUserKey string    `gorm:"size:64" json:"bind_user_key"`
	State       int       `gorm:"size:2" json:"state"`
	LoginAt     time.Time `json:"login_at"`
	LastPingAt  time.Time `json:"last_ping_at"`
}
