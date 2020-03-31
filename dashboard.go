package main

import (
	"Treasuro/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type Data struct {
	User     string
	Level    int
	Score    int
	Attempts int
	Question string
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	fmt.Println("token", tokenString)

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
			w.Write([]byte(`{"error": "Token not verified"}`))
		}
		return []byte("idgafaboutthingsanymore"), nil
	})
	var _, id string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		_ = claims["name"].(string)
		id = claims["id"].(string)
	}
	user := database.Finddb(cl1, id)
	if user.Attempts==5{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Sorry, you have been disqualified"}`))
	}else{
		ques := database.FindQuestion(cl3, user.Level)
		data := Data{
			User:     user.Username,
			Level:    user.Level,
			Score:    user.Score,
			Attempts: user.Attempts,
			Question: ques.Question,
		}
		json.NewEncoder(w).Encode(data)
		w.WriteHeader(http.StatusOK)
	}
}

func submit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	fmt.Println("token", tokenString)

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
			w.Write([]byte(`{"error": "Token not verified"}`))
		}
		return []byte("idgafaboutthingsanymore"), nil
	})
	var _, id string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		_ = claims["name"].(string)
		id = claims["id"].(string)
	}
	user := database.Finddb(cl1, id)
	var Answer string
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &Answer)
	if err != nil {
		ques := database.FindQuestion(cl3, user.Level)
		if ques.Answer == Answer {
			database.UpdateScore(cl1, id, 10)
		} else {
			database.UpdateAttempts(cl1, id)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"successfull": "updated score"}`))
	}else{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "password do not match"}`))
	}
    
}
