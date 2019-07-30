package dto

import "time"

type EnvironmentComponentReq struct {
	EnvironmentComponentKey          string `json:"environmentComponentKey"`
	EnvironmentComponentName         string `json:"environmentComponentName"`
	EnvironmentComponentDescription  string `json:"environmentComponentDescription"`
	EnvironmentComponentVersionCount int    `json:"environmentComonentVersionCount`
	CreatedAt                       time.Time
	UpdatedAt                       time.Time
}