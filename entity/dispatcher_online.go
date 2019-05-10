package entity

type DispatcherOnline struct {
	OnlineKey                 string            `gorm:"primary_key;size:64" json:"onlineKey"`
	RelationMachineSign       string            `gorm:"size:64;"json:"relationMachineSign"`
	RelationDispatcherMachine DispatcherMachine `gorm:"ForeignKey:MachineSign;AssociationForeignKey:RelationMachineSign;" json:"relationDispatcherMachine"`
	IpAddress                 string            `gorm:"size:32" json:"ipAddress"`
	BindServerMachineSign     string            `gorm:"size:64" json:"bindServerMachineSign"`
}
