package model

type LrConfig struct {
	DbType        string `json:"db_type"`
	DbUrl         string `json:"db_url"`
	WorkSpacePath string `json:"work_space_path"`
}
