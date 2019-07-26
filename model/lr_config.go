package model

type LrServerConfig struct {
	Host                      string            `json:"host"`
	Port                      int               `json:"port"`
	DbType                    string            `json:"dbType"`
	DbUrl                     string            `json:"dbUrl"`
	DbTablePrefix             string            `json:"dbTablePrefix"`
	LoginAuthLength           int               `json:"loginAuthLength"` // unit: minutes
	WorkSpacePath             string            `json:"workSpacePath"`
	DebugMode                 bool              `json:"debugMode"`
	SecretHmacKeyword         string            `json:"secretHmacKeyword"`
	ClusterNodeActiveInterval int               `json:"clusterNodeActiveInterval"` // unit: seconds
	FileResourceType          string            `json:"fileResourceType"`
	FileResourceConfig        map[string]string `json:"fileResourceConfig"`
	GitType                   string            `json:"gitType"`
	GitConfig                 map[string]string `json:"gitConfig"`
}
