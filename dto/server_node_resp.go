package dto

import "lemon-robot-server/entity"

type ServerNodeResp struct {
	NodeInfo    entity.ServerNode `json:"node_info"`
	ActiveState bool              `json:"active_state"`
}
