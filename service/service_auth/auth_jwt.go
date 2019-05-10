package service_auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"lemon-robot-golang-commons/utils/lrustring"
	"lemon-robot-server/dao/dao_user"
	"lemon-robot-server/entity"
	"lemon-robot-server/model"
	"lemon-robot-server/sysinfo"
	"time"
)

func GenerateJwtTokenStr(userKey string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, generateJwtPayload(userKey))
	tokenString, _ := token.SignedString(sysinfo.GetHmacKeyBytes())
	return tokenString
}

func generateJwtPayload(userKey string) model.LrJwtPayload {
	expireDur, _ := time.ParseDuration(fmt.Sprintf("%dm", sysinfo.LrServerConfig().LoginAuthLength))
	return model.LrJwtPayload{
		Id:        lrustring.Uuid(),
		Issuer:    userKey,
		IssuedAt:  time.Now().Unix(),
		NotBefore: time.Now().Unix(),
		ExpiresAt: time.Now().Add(expireDur).Unix(),
		Audience:  userKey,
		Subject:   sysinfo.AppName(),
	}
}

func CheckToken(jwtTokenStr string) bool {
	jwtToken, err := jwt.Parse(jwtTokenStr, func(token *jwt.Token) (i interface{}, e error) {
		return sysinfo.GetHmacKeyBytes(), nil
	})
	if jwtToken == nil {
		return false
	}
	userKey := jwtToken.Claims.(jwt.MapClaims)["iss"]
	user := dao_user.FirstByExample(&entity.User{UserKey: userKey.(string)})
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
