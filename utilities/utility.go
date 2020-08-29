package utilities

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(name, id string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
		"id":   id,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("blockkey")))
	fmt.Println("generated",tokenString)
	if err == nil {
		return tokenString
	}
	return ""
}

func RandomScore() int{
	rand.Seed(time.Now().UnixNano())
	min := -10
	max := 10
	return (rand.Intn(max-min+1) + min)
}
