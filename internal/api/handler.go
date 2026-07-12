package api

import (
	"encoding/json"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"status": "ok",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
