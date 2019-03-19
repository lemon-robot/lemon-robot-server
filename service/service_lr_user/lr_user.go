package service_lr_user

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/utils/lrustring"
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
	"lemon-robot-server/sysinfo"
)

var HashObj hash.Hash

func GenerateLrUser(number, password string) entity.LrUser {
	lrUserEntity := entity.LrUser{}
	lrUserEntity.LrUserKey = lrustring.Uuid()
	lrUserEntity.UserNumber = number
	lrUserEntity.PasswordSecret = CalculatePasswordSecret(password)
	db.Db().Create(&lrUserEntity)
	return lrUserEntity
}

func CheckLrc(number, password string) (bool, entity.LrUser) {
	lrcEntity := entity.LrUser{}
	db.Db().First(&lrcEntity, &entity.LrUser{UserNumber: number})
	if lrcEntity.ID == 0 {
		return false, lrcEntity
	}
	return lrcEntity.PasswordSecret == CalculatePasswordSecret(password), lrcEntity
}

func CalculatePasswordSecret(password string) string {
	if HashObj == nil {
		HashObj = hmac.New(sha256.New, sysinfo.GetHmacKeyBytes())
	}
	HashObj.Reset()
	io.WriteString(HashObj, password)
	return fmt.Sprintf("%x", HashObj.Sum(nil))
}

// lrc self repair,
// If the LRC table in the database is empty,
// then an LRC is automatically created randomly and the information is displayed in the console.
func SelfRepair() {
	logger.Info("Start self-repair: LR User")
	//count := db.Db().Count(&entity.Lrc{})
	var count int
	db.Db().Model(&entity.LrUser{}).Count(&count)
	if count == 0 {
		numberNew := "lemonrobot"
		passwordNew := lrustring.Uuid()
		GenerateLrUser(numberNew, passwordNew)
		logger.Warn("There are no users in the system. Now the system will automatically generate a user.")
		logger.Warn("Number: " + numberNew)
		logger.Warn("Password: " + passwordNew)
	}
	logger.Info("Self-repair completed: LR User")
}
