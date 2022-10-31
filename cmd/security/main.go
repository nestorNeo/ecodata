package main

// this is a security util for generating tokens
import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/nestorneo/ecodata/security"
)

var (
	key   string
	email string
	year  int
	month int
	day   int
)

func init() {
	flag.StringVar(&email, "email", "", "user email")
	flag.IntVar(&year, "year", 1996, "birth year")
	flag.IntVar(&month, "month", 3, "birth month")
	flag.IntVar(&day, "day", 22, "birth day")
	// signing key
	flag.StringVar(&key, "key", "", "signing key")

}

func main() {

	flag.Parse()

	if len(email) == 0 || len(key) == 0 {
		log.Fatalln("invalid data, do help")
	}

	testUser := security.EcoUser{
		Email:    email,
		Birthday: time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).Unix(),
		StandardClaims: jwt.StandardClaims{
			// expires in 3 months
			ExpiresAt: time.Now().AddDate(0, 3, 0).Unix(),
			Issuer:    "Friends of Go",
		},
	}

	validator, err := security.NewValidator(key, key+".pub")

	if err != nil {
		log.Fatalln("error reading private key", err)
	}

	token, err := validator.SignAndCreate(&testUser)
	if err != nil {
		log.Fatalln("cannot SIGN", err)
	}

	fmt.Println(token)
}
