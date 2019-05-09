package model

type LrServerConfig struct {
	Host                      string            `json:"host"`
	Port                      int               `json:"port"`
	DbType                    string            `json:"dbType"`
	DbUrl                     string            `json:"dbUrl"`
	DbTablePrefix             string            `json:"dbTablePrefix"`
	WorkSpacePath             string            `json:"workSpacePath"`
	DebugMode                 bool              `json:"debugMode"`
	SecretHmacKeyword         string            `json:"secretHmacKeyword"`
	ClusterNodeActiveInterval int               `json:"clusterNodeActiveInterval"` // unit: seconds
	GitType                   string            `json:"gitType"`
	GitConfig                 map[string]string `json:"gitConfig"`
}
