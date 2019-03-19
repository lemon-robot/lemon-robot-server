package service_auth

import (
	"github.com/dgrijalva/jwt-go"
	"lemon-robot-golang-commons/utils/lrustring"
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
	return model.LrJwtPayload{
		Id:        lrustring.Uuid(),
		Issuer:    userKey,
		IssuedAt:  time.Now().Unix(),
		NotBefore: time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
		Audience:  userKey,
		Subject:   sysinfo.AppName(),
	}
}
