package main

import (
	"Treasuro/database"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var cl1, cl2 *mongo.Collection

func init(){
 cl1,cl2:=database.Createdb()
 ok:=database.InsertUUIDindb(cl2)
 if ok{
	 fmt.Print("uuid's generated")
 }
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", register).Methods("POST")
	r.HandleFunc("/signup",signup).Methods("POST","GET")
	// r.HandleFunc("/login", login).Methods("POST")
	// r.HandleFunc("/dashboard", dashboard).Methods("GET", "POST")
	http.Handle("/", r)
	http.ListenAndServe(":80", nil)
}
