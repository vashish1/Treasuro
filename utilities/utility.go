package utilities

import "github.com/dgrijalva/jwt-go"

func GenerateToken(name,id string)string{
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":  name,
		"id": id,
	})

	tokenString, err := token.SignedString([]byte("idgafaboutthingsanymore"))
	if err!=nil {
		return tokenString
	}
	return ""
}