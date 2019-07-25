package dto

import "time"

type EnvironmentComponentReq struct {
	EnvironmentComponentKey         string `json:"environmentComponentKey"`
	EnvironmentComponentName        string `json:"environmentComponentName"`
	EnvironmentComponentDescription string `json:"environmentComponentDescription"`
	CreatedAt                       time.Time
	UpdatedAt                       time.Time
}
