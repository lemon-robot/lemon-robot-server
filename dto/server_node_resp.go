package dto

import "lemon-robot-server/entity"

type ServerNodeResp struct {
	NodeInfo    entity.ServerNode `json:"nodeInfo"`
	ActiveState bool              `json:"activeState"`
}
