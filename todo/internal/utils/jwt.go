package utils

import (
	"time"
	"todo/configs"
	"todo/internal/features/users"

	"github.com/golang-jwt/jwt/v5"
)

type TokenUtilInterface interface {
	GenerateToken(users.User) (string, error)
	DecodeToken(*jwt.Token) uint
}

type tokenUtil struct{}

func NewTokenUtil() TokenUtilInterface {
	return &tokenUtil{}
}

func (tu *tokenUtil) GenerateToken(LoginData users.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = LoginData.ID
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	process := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := process.SignedString([]byte(configs.ImportSetting().Passkey))

	if err != nil {
		return "", err
	}

	return result, nil
}

func (tu *tokenUtil) DecodeToken(token *jwt.Token) uint {
	claims := token.Claims.(jwt.MapClaims)

	return uint(claims["id"].(float64))
}
