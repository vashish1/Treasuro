package main

import (
	"Treasuro/database"
	"encoding/json"
	"net/http"
)

func leaderboard(w http.ResponseWriter, r *http.Request) {
	list := database.Leaderboard(cl1)
	json.NewEncoder(w).Encode(list)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"success": "fetched data successfully"}`))
}
