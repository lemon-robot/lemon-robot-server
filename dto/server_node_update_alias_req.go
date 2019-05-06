package dto

type ServerNodeUpdateAliasReq struct {
	MachineSign string `json:"machine_sign"`
	Alias       string `json:"alias"`
}
