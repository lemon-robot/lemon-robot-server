package entity

type DispatcherOnline struct {
	OnlineKey                 string            `gorm:"primary_key;size:64" json:"onlineKey"`
	RelationDispatcherMachine DispatcherMachine `gorm:"ForeignKey:RelationMachineSign;AssociationForeignKey:MachineSign" json:"relationDispatcherMachine"`
	RelationMachineSign       string            `json:"relationMachineSign"`
	IpAddress                 string            `gorm:"size:32" json:"ipAddress"`
	BindServerMachineSign     string            `gorm:"size:64" json:"bindServerMachineSign"`
}
