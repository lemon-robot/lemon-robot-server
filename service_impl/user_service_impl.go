package service_impl

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/utils/lru_string"
	"lemon-robot-server/dao"
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
)

type UserServiceImpl struct {
	hashObj hash.Hash
	userDao *dao.UserDao
}

func NewUserServiceImpl(passwordSecretKey string) *UserServiceImpl {
	return &UserServiceImpl{
		hashObj: hmac.New(sha256.New, []byte(passwordSecretKey)),
		userDao: dao.NewUserDao(),
	}
}

func (i *UserServiceImpl) Create(number, password string) (error, entity.User) {
	userEntity := entity.User{
		UserKey:        lru_string.GetInstance().Uuid(true),
		PasswordSecret: i.calculateSecretPassword(password),
		UserNumber:     number}
	result := db.Db().Create(&userEntity)
	return result.Error, userEntity
}

func (i *UserServiceImpl) CountByNumber(number string) int {
	return i.userDao.CountByUserExample(&entity.User{UserNumber: number})
}

func (i *UserServiceImpl) CheckPassword(number, password string) (bool, entity.User) {
	userEntity := i.userDao.FirstByExample(&entity.User{UserNumber: number})
	if userEntity.UserKey == "" {
		return false, userEntity
	}
	return userEntity.PasswordSecret == i.calculateSecretPassword(password), userEntity
}

func (i *UserServiceImpl) calculateSecretPassword(password string) string {
	i.hashObj.Reset()
	_, err := io.WriteString(i.hashObj, password)
	if err != nil {
		logger.Error("Error when calculate password secret", err)
	}
	return fmt.Sprintf("%x", i.hashObj.Sum(nil))
}

// lr user self repair,
// If the User table in the database is empty,
// then an User is automatically created randomly and the information is displayed in the console.
func (i *UserServiceImpl) SelfRepair() {
	logger.Info("Start self-repair: User")
	var count int
	db.Db().Model(&entity.User{}).Count(&count)
	if count == 0 {
		numberNew := "lemonrobot"
		passwordNew := lru_string.GetInstance().Uuid(true)
		err, _ := i.Create(numberNew, passwordNew)
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
