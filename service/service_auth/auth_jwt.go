package service_auth

import (
	"github.com/dgrijalva/jwt-go"
	"lemon-robot-golang-commons/utils/lrustring"
	"lemon-robot-server/model"
	"lemon-robot-server/sysinfo"
	"time"
)

func GenerateJwtTokenStr(lrct string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, generateJwtPayload(lrct))
	tokenString, _ := token.SignedString(sysinfo.GetHmacKeyBytes())
	return tokenString
}

func generateJwtPayload(lrct string) model.LrJwtPayload {
	return model.LrJwtPayload{
		Id:        lrustring.Uuid(),
		Issuer:    lrct,
		IssuedAt:  time.Now().Unix(),
		NotBefore: time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
		Audience:  lrct,
		Subject:   sysinfo.AppName(),
	}
}
