package service_user

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/utils/lrustring"
	"lemon-robot-server/dao/dao_user"
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
	"lemon-robot-server/sysinfo"
)

var HashObj hash.Hash

func CreateUser(number, password string) (error, entity.User) {
	userEntity := entity.User{}
	userEntity.UserKey = lrustring.Uuid()
	userEntity.UserNumber = number
	userEntity.PasswordSecret = CalculatePasswordSecret(password)
	result := db.Db().Create(&userEntity)
	return result.Error, userEntity
}

func CheckUser(number, password string) (bool, entity.User) {
	userEntity := dao_user.FirstByExample(&entity.User{UserNumber: number})
	if userEntity.UserKey == "" {
		return false, userEntity
	}
	return userEntity.PasswordSecret == CalculatePasswordSecret(password), userEntity
}

func CalculatePasswordSecret(password string) string {
	if HashObj == nil {
		HashObj = hmac.New(sha256.New, sysinfo.GetHmacKeyBytes())
	}
	HashObj.Reset()
	_, err := io.WriteString(HashObj, password)
	if err != nil {
		logger.Error("Error when calculate password secret", err)
	}
	return fmt.Sprintf("%x", HashObj.Sum(nil))
}

// lr user self repair,
// If the User table in the database is empty,
// then an User is automatically created randomly and the information is displayed in the console.
func SelfRepair() {
	logger.Info("Start self-repair: User")
	var count int
	db.Db().Model(&entity.User{}).Count(&count)
	if count == 0 {
		numberNew := "lemonrobot"
		passwordNew := lrustring.Uuid()
		err, _ := CreateUser(numberNew, passwordNew)
		if err != nil {
			logger.Error("Error repair User, Can not create user", err)
			return
		}
		logger.Warn("There are no users in the system. Now the system will automatically generate a user.")
		logger.Warn("Number: " + numberNew)
		logger.Warn("Password: " + passwordNew)
	}
	logger.Info("Self-repair completed: User")
}
