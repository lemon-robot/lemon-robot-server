package model

type LrConfig struct {
	DbType            string `json:"db_type"`
	DbUrl             string `json:"db_url"`
	WorkSpacePath     string `json:"work_space_path"`
	DebugMode         bool   `json:"debug_mode"`
	SecretHmacKeyword string `json:"secret_hmac_keyword"`
}
