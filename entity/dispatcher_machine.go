package entity

type DispatcherMachine struct {
	MachineSign       string          `gorm:"primary_key;size:64" json:"machine_sign"`
	CpuArch           string          `gorm:"size:32" json:"cpu_arch"`
	OperateSystem     string          `gorm:"size:32" json:"operate_system"`
	DispatcherVersion string          `gorm:"size:32" json:"dispatcher_version"`
	Tags              []DispatcherTag `gorm:"many2many:dispatcher_tag_relation;ForeignKey:MachineSign;AssociationForeignKey:TagKey" json:"tags"`
	Alias             string          `gorm:"size:64" json:"alias"`
}
