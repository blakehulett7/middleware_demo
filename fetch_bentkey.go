package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func fetch_bentkey(w http.ResponseWriter, r *http.Request) {
	auth_header := r.Header.Get("Authorization")
	token := strings.Split(auth_header, " ")[1]

	response := map[string]string{
		"token": token,
	}
	payload, _ := json.Marshal(response)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(payload)
}
