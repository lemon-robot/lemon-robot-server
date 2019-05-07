package entity

type DispatcherMachine struct {
	MachineSign       string          `gorm:"primary_key;size:64" json:"machineSign"`
	CpuArch           string          `gorm:"size:32" json:"cpuArch"`
	OperateSystem     string          `gorm:"size:32" json:"operateSystem"`
	DispatcherVersion string          `gorm:"size:32" json:"dispatcherVersion"`
	Tags              []DispatcherTag `gorm:"many2many:dispatcher_tag_relation;ForeignKey:MachineSign;AssociationForeignKey:TagKey" json:"tags"`
	Alias             string          `gorm:"size:64" json:"alias"`
}
