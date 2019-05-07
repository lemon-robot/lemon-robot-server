package entity

import "time"

type ServerNode struct {
	MachineSign       string             `gorm:"primary_key;size:64" json:"machineSign"`
	CpuArch           string             `gorm:"size:32" json:"cpuArch"`
	OperateSystem     string             `gorm:"size:32" json:"operateSystem"`
	ServerVersion     string             `gorm:"size:32" json:"serverVersion"`
	Alias             string             `gorm:"size:64" json:"alias"`
	StartAt           time.Time          `json:"startAt"`
	ActiveTime        time.Time          `json:"activeTime"`
	OnlineDispatchers []DispatcherOnline `gorm:"ForeignKey:MachineSign;AssociationForeignKey:BindServerMachineSign" json:"onlineDispatchers"`
}
