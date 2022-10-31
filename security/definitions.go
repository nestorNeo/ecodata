package security

import (
	"crypto/rsa"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

type Validator struct {
	priv *rsa.PrivateKey
	pub  *rsa.PublicKey
}

func NewValidator(keyPriv, keyPub string) (*Validator, error) {

	privData, err := os.ReadFile(keyPriv)

	if err != nil {
		return nil, err
	}

	pubData, err := os.ReadFile(keyPub)
	if err != nil {
		return nil, err
	}

	priv, err := jwt.ParseRSAPrivateKeyFromPEM(privData)

	if err != nil {
		return nil, err
	}

	pub, err := jwt.ParseRSAPublicKeyFromPEM(pubData)

	if err != nil {
		return nil, err
	}

	return &Validator{
		priv,
		pub,
	}, nil
}

type EcoUser struct {
	Email    string `json:"email"`
	Birthday int64  `json:"birthday"`
	jwt.StandardClaims
}

func (val *Validator) SignAndCreate(user *EcoUser) (string, error) {
	// tokens need to be updated for security dummy function
	return jwt.NewWithClaims(jwt.SigningMethodRS256, *user).SignedString(val.priv)
}

func (val *Validator) IsValid(token string) (jwt.MapClaims, error) {
	// tokens need to be updated for security dummy function

	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}
		return val.pub, nil
	})

	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, fmt.Errorf("validate: invalid")
	}

	return claims, nil
}
