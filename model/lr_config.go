package model

type LrConfig struct {
	Host                          string            `json:"host"`
	Port                          int               `json:"port"`
	DbType                        string            `json:"db_type"`
	DbUrl                         string            `json:"db_url"`
	DbTablePrefix                 string            `json:"db_table_prefix"`
	WorkSpacePath                 string            `json:"work_space_path"`
	DebugMode                     bool              `json:"debug_mode"`
	SecretHmacKeyword             string            `json:"secret_hmac_keyword"`
	ClusterNodeActiveInterval     int               `json:"cluster_node_active_interval"`      // unit: seconds
	ClusterNodeActiveScanInterval int               `json:"cluster_node_active_scan_interval"` // unit: seconds
	GitType                       string            `json:"git_type"`
	GitConfig                     map[string]string `json:"git_config"`
}
