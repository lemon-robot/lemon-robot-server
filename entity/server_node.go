package entity

import "time"

type ServerNode struct {
	MachineSign       string             `gorm:"primary_key;size:64" json:"machine_sign"`
	CpuArch           string             `gorm:"size:32" json:"cpu_arch"`
	OperateSystem     string             `gorm:"size:32" json:"operate_system"`
	ServerVersion     string             `gorm:"size:32" json:"server_version"`
	Alias             string             `gorm:"size:64" json:"alias"`
	StartAt           time.Time          `json:"start_at"`
	ActiveTime        time.Time          `json:"active_time"`
	OnlineDispatchers []DispatcherOnline `gorm:"ForeignKey:MachineSign;AssociationForeignKey:BindServerMachineSign" json:"online_dispatchers"`
}
