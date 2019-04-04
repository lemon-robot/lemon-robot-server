package entity

type DispatcherOnline struct {
	OnlineKey                 string            `gorm:"primary_key;size:64" json:"online_key"`
	RelationDispatcherMachine DispatcherMachine `gorm:"ForeignKey:RelationMachineCode;AssociationForeignKey:MachineCode" json:"relation_dispatcher_machine"`
	RelationMachineCode       string            `json:"relation_machine_code"`
	IpAddress                 string            `gorm:"size:32" json:"ip_address"`
	BindServerMachineCode     string            `gorm:"size:64" json:"bind_server_machine_code"`
}
