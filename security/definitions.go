package security

import (
	"github.com/golang-jwt/jwt/v4"
)

type EcoUser struct {
	Email    string `json:"email"`
	Birthday int64  `json:"birthday"`
	jwt.StandardClaims
}

func SignAndCreate(user *EcoUser, key []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, *user)
	return token.SignedString(key)
}
