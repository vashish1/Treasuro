package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/vashish1/Treasuro/database"

	"github.com/dgrijalva/jwt-go"
)

type logn struct {
	Username string
	Password string
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var result Register
	var user logn
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}
	ok := database.FindUser(cl1, user.Username, user.Password)
	if ok {
		u := database.Finddb(cl1, user.Username)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"name": u.Name,
			"id":   u.UUID,
		})

		tokenString, err := token.SignedString([]byte(secret))

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "error in token string"}`))
			return
		}
	     	result.Success=true
			result.Token=tokenString
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(result)
			return
		}
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Invalid Credentials"}`))
}

