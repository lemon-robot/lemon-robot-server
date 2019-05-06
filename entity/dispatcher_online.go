package entity

type DispatcherOnline struct {
	OnlineKey                 string            `gorm:"primary_key;size:64" json:"online_key"`
	RelationDispatcherMachine DispatcherMachine `gorm:"ForeignKey:RelationMachineSign;AssociationForeignKey:MachineSign" json:"relation_dispatcher_machine"`
	RelationMachineSign       string            `json:"relation_machine_sign"`
	IpAddress                 string            `gorm:"size:32" json:"ip_address"`
	BindServerMachineSign     string            `gorm:"size:64" json:"bind_server_machine_sign"`
}
