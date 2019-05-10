package service

import "lemon-robot-server/entity"

type UserService interface {
	Create(number, password string) (error, entity.User)
	CountByNumber(number string) int
	CheckPassword(number, password string) (bool, entity.User)
}
