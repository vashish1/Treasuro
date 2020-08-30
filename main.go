package main

import (
	"fmt"

	"github.com/vashish1/Treasuro/database"

	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var cl1, cl2, cl3 *mongo.Collection

func init() {
	fmt.Println("secret", secret)
	cl1, cl2, cl3 = database.Createdb()
}

func main() {
	r := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	r.HandleFunc("/register", register).Methods("POST")
	r.HandleFunc("/signup", signup).Methods("POST", "GET")
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/dashboard", dashboard).Methods("GET")
	r.HandleFunc("/submit", submit).Methods("POST")
	r.HandleFunc("/leaderboard", leaderboard).Methods("GET")
	http.Handle("/", handlers.CORS(headers, methods, origins)(r))
	http.ListenAndServe(":80", nil)
}
