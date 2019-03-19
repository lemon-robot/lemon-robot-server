package model

type LrConfig struct {
	DbType            string            `json:"db_type"`
	DbUrl             string            `json:"db_url"`
	DbTablePrefix     string            `json:"db_table_prefix"`
	WorkSpacePath     string            `json:"work_space_path"`
	DebugMode         bool              `json:"debug_mode"`
	SecretHmacKeyword string            `json:"secret_hmac_keyword"`
	GitType           string            `json:"git_type"`
	GitConfig         map[string]string `json:"git_config"`
}
