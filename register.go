package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/vashish1/Treasuro/database"
)

var secret=os.Getenv("blockkey")

type regis struct {
	Name string
	Uid  string
}

type mockSignup struct {
	Username  string
	Number  string
	Email     string
	Password  string
	Cpassword string
}

type Register struct {
	Success bool   `json:"success,omitempty"`
	Token   string `json:"token,omitempty"`
}

func register(w http.ResponseWriter, r *http.Request) {
	var input regis
	var result Register
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}
	ok := database.UuidExists(cl2, input.Uid)
	if ok {

		okk,token := database.RegisterUser(cl1, input.Name, input.Uid)
		if okk {
			w.WriteHeader(http.StatusOK)
			result.Token = token
			result.Success = true
			json.NewEncoder(w).Encode(result)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "uuid already registered"}`))
		return
	} 
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "uuid do not exist"}`))
}

func signup(w http.ResponseWriter, r *http.Request) {
	var test mockSignup
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	fmt.Println("token",tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(secret), nil
	})
	fmt.Println("err",err)
	var _, id string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id = claims["id"].(string)
		_ = claims["name"].(string)
	}
	fmt.Println("////",id)
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &test)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}
	if ok := database.CheckUsername(cl1, test.Username); !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Username already exists"}`))
		return
	}
	if test.Password == test.Cpassword {
		p := database.UpdateUserCreds(cl1, id, test.Username, test.Number, test.Email, test.Password)
		if p {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"success": "true"}`))
			return
		}
		    w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"success": "false"}`))
			return
	} 
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Password do not match"}`))
}
