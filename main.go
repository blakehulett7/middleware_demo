package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Jesus is Lord!")

	router := http.NewServeMux()
	router.HandleFunc("/middleware-demo", handle_test)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}

func handle_test(w http.ResponseWriter, r *http.Request) {
	res := struct {
		Res string `json:"res"`
	}{"Jesus is Lord!"}
	data, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}
