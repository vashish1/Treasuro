package utilities

import (
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(name, id string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
		"id":   id,
	})

	tokenString, err := token.SignedString([]byte("idgafaboutthingsanymore"))
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
