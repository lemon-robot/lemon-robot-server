package dto

type EnvironmentComponentReq struct {
	EnvironmentComponentKey          string `json:"environmentComponentKey"`
	EnvironmentComponentName         string `json:"environmentComponentName"`
	EnvironmentComponentDescription  string `json:"environmentComponentDescription"`
	EnvironmentComponentVersionCount int    `json:"environmentComonentVersionCount`
	IconFileResourceKey              string `gorm:"size:64" json:"iconFileResourceKey"`
	IconFileResourceUrl              string `json:"iconFileResourceUrl"`
}
