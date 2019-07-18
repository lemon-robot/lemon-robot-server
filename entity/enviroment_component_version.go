package entity

import "time"

type EnvironmentComponentVersion struct {
	ECVersionKey                  string               `gorm:"size:64;unique_index" json:"ecVersionKey"`
	BelongEnvironmentComponent    EnvironmentComponent `gorm:"ForeignKey:BelongEnvironmentComponentKey;AssociationForeignKey:EnvironmentComponentKey" json:"belongEnvironmentComponent"`
	BelongEnvironmentComponentKey string               `gorm:"primary_key;size:64" json:"belongEnvironmentComponentKey"`
	ECVersionNumber               int                  `gorm:"primary_key;size:4" json:"ecVersionNumber"`
	ECVersionName                 string               `gorm:"size:128" json:"ecVersionName"`
	ECVersionDescription          string               `gorm:"size:10240" json:"ecVersionDescription"`
	StateCheckScript              string               `gorm:"type:longtext" json:"stateCheckScript"`
	InstallScript                 string               `gorm:"type:longtext" json:"installScript"`
	UninstallScript               string               `gorm:"type:longtext" json:"uninstallScript"`
	ProgramFileResource           FileResource         `gorm:"ForeignKey:ProgramFileResourceKey;AssociationForeignKey:FileResourceKey" json:"programFileResource"`
	ProgramFileResourceKey        string
	WhereToInstall                string `gorm:"size:32" json:"whereToInstall"`
	CreatedAt                     time.Time
	UpdatedAt                     time.Time
	DeletedAt                     *time.Time `sql:"index"`
}
