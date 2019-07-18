package service_impl

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"lemon-robot-golang-commons/utils/lru_string"
	"lemon-robot-server/dao"
	"lemon-robot-server/entity"
	"lemon-robot-server/model"
	"lemon-robot-server/sysinfo"
	"time"
)

type AuthServiceImpl struct {
	userDao *dao.UserDao
}

func NewAuthServiceImpl() *AuthServiceImpl {
	return &AuthServiceImpl{
		userDao: dao.NewUserDao(),
	}
}

func (i *AuthServiceImpl) GenerateJwtTokenStr(userKey string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, i.generateJwtPayload(userKey))
	tokenString, _ := token.SignedString(sysinfo.GetHmacKeyBytes())
	return tokenString
}

func (i *AuthServiceImpl) generateJwtPayload(userKey string) model.LrJwtPayload {
	expireDur, _ := time.ParseDuration(fmt.Sprintf("%dm", sysinfo.LrServerConfig().LoginAuthLength))
	return model.LrJwtPayload{
		Id:        lru_string.GetInstance().Uuid(true),
		Issuer:    userKey,
		IssuedAt:  time.Now().Unix(),
		NotBefore: time.Now().Unix(),
		ExpiresAt: time.Now().Add(expireDur).Unix(),
		Audience:  userKey,
		Subject:   sysinfo.AppName(),
	}
}

func (i *AuthServiceImpl) CheckToken(jwtTokenStr string) bool {
	jwtToken, err := jwt.Parse(jwtTokenStr, func(token *jwt.Token) (i interface{}, e error) {
		return sysinfo.GetHmacKeyBytes(), nil
	})
	if jwtToken == nil {
		return false
	}
	userKey := jwtToken.Claims.(jwt.MapClaims)["iss"]
	user := i.userDao.FirstByExample(&entity.User{UserKey: userKey.(string)})
	// user not found or have error
	if user.UserKey == "" || err != nil {
		return false
	}
	if _, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		return true
	} else {
		return false
	}
}
