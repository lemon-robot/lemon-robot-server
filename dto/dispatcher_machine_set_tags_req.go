package dto

type DispatcherMachineSetTagsReq struct {
	MachineSign string   `json:"machineSign"`
	TagKeys     []string `json:"tagKeys"`
}
