package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Jesus is Lord!")

	router := http.NewServeMux()

	router.HandleFunc("/middleware-test", handle_test)
	router.HandleFunc("/middleware-test2", specific_middleware(handle_test))
	router.HandleFunc("/middleware-test3", handle_test)

	handler := general_middleware(router)
	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	server.ListenAndServe()
}

func handle_test(w http.ResponseWriter, r *http.Request) {
	res := struct {
		Res string `json:"res"`
	}{"Jesus is Lord"}

	data, _ := json.Marshal(res)

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func general_middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("do general middleware thing")
		h.ServeHTTP(w, r)
	})
}

func specific_middleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("do specific middleware thing")
		h(w, r)
	}
}
