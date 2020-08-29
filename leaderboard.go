package main

import (
	"encoding/json"
	"net/http"

	"github.com/vashish1/Treasuro/database"
)

func leaderboard(w http.ResponseWriter, r *http.Request) {
	list := database.Leaderboard(cl1)
	json.NewEncoder(w).Encode(list)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"success": "fetched data successfully"}`))
}
