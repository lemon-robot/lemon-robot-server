package entity

type Config struct {
	ConfigKey   string `gorm:"primary_key;size:256" json:"config_key"`
	ConfigValue string `gorm:"type:longtext" json:"config_value"`
}
