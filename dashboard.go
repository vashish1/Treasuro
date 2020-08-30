package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/vashish1/Treasuro/database"

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
		}
		return []byte(secret), nil
	})
	var _, id string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		_ = claims["name"].(string)
		id = claims["id"].(string)
	}
	user := database.Finddb(cl1, id)
	if user.Attempts == 5 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Sorry, you have been disqualified"}`))
		return
	}
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

func submit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	fmt.Println("token", tokenString)
	var resp Register
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(secret), nil
	})
	var _, id string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		_ = claims["name"].(string)
		id = claims["id"].(string)
	}
	var Answer database.Response
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &Answer)
	if err == nil {
		score := database.CheckAnswer(cl3, Answer)
		if score == 0 {
			database.UpdateAttempts(cl1,id)
			w.WriteHeader(http.StatusNotAcceptable)
			w.Write([]byte(`{"success": false}`))
			return
		}
		if score==10{
			database.FixAttempts(cl1,id)
		}
		database.UpdateScore(cl1, id, score)
		resp.Success = true
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`{"error": "Body not parsed"}`))
}
