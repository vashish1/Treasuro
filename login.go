package main

import (
	"Treasuro/database"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type logn struct {
	Username string
	Password string
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// var result database.User
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

		tokenString, err := token.SignedString([]byte("idgafaboutthingsanymore"))

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "error in token string"}`))
			return
		}
		tkn := database.UpdateToken(cl1, u.Email, tokenString)
		if tkn {
			json.NewEncoder(w).Encode(tokenString)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"success": "created token successfully"}`))
		} else {
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"error": "token not created"}`))
		}
	}
}
