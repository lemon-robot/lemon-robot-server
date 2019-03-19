package model

type ReqUserCreate struct {
	Number   string `json:"number"`
	Password string `json:"password"`
}
