package entity

type Config struct {
	ConfigKey   string `gorm:"primary_key;size:256" json:"configKey"`
	ConfigValue string `gorm:"type:longtext" json:"configValue"`
}
