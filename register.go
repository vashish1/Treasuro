package main

import (
	"Treasuro/database"
	"Treasuro/utilities"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
	"strings"
)

type regis struct {
	Name string
	Uid  string
}

type mockSignup struct {
	Username string
	PhNumber string
	Email     string 
	Password  string 
	Cpassword string 
}

func register(w http.ResponseWriter, r *http.Request) {
	var test regis
	w.Header().Set("Content-Type", "application/json")
	var user database.User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &test)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}
	ok:=database.UuidExists(cl2,test.Uid)
	if ok{
		user.Name=test.Name
		user.UUID=test.Uid
		user.Token=utilities.GenerateToken(user.Name,user.UUID)
		okk:=database.Insertintouserdb(cl1,user)
		if okk{
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(user.Token)
			w.Write([]byte(`{"success": "created token successfully"}`))
		    w.Write([]byte(`{"successfull": "user created"}`))
		}

	}else{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "uuid do not exist"}`))
	}

}

func signup(w http.ResponseWriter, r *http.Request){
    var test mockSignup
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
		 _= claims["name"].(string)
		id = claims["id"].(string)
	}
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &test)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}
    if ok:=database.CheckUsername(cl1,test.Username);!ok{
        w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Username already exists"}`))
		return
	}
	if(test.Password==test.Cpassword){
		p:=database.UpdateUserCreds(cl1,id,test.Username,test.PhNumber,test.Email,test.Password)
		if p{
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"successfull": "updated credentials"}`))
		}
	}else{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "password do not match"}`))
	}
}