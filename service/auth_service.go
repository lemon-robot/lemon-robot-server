package service

type AuthService interface {
	GenerateJwtTokenStr(userKey string) string
	CheckToken(jwtTokenStr string) bool
}
