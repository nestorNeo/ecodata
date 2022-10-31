package main

// this is a security util for generating tokens
import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/nestorneo/ecodata/security"
)

func main() {
	testUser := security.EcoUser{
		Email:    "test@gdl.cinvestav.mx",
		Birthday: time.Date(2019, 01, 01, 0, 0, 0, 0, time.UTC).Unix(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "Friends of Go",
		},
	}

	// NOT SECURE THIS IS A TEST KEY
	key := "oW1C48WRN^vd"

	if len(os.Getenv("ECODATA_SIGN_KEY")) > 0 {
		key = os.Getenv("ECODATA_SIGN_KEY")
	}

	token, err := security.SignAndCreate(&testUser, []byte(key))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(token)
}
