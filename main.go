package main

import (
	"Treasuro/database"

	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var cl1, cl2,cl3  *mongo.Collection

func init(){
 cl1,cl2,cl3=database.Createdb()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", register).Methods("POST")
	r.HandleFunc("/signup",signup).Methods("POST","GET")
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/dashboard", dashboard).Methods("GET")
	r.HandleFunc("/submit",submit).Methods("POST")
    r.HandleFunc("/leaderboard",leaderboard).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":80", nil)
}
