package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Jesus is Lord!")

	router := http.NewServeMux()

	router.HandleFunc("/middleware-test", handle_test)
	router.HandleFunc("/middleware-test2", specific_middleware(handle_test))

	router.HandleFunc("/protected/guarded-endpoint", guardian_middleware(handle_test))
	router.HandleFunc("/protected/guarded-endpoint2", guardian_middleware(handle_test))

	router.Handle("/prefix/", prefix_middleware(handle_test))

	handler := general_middleware(router)
	server := http.Server{
		Addr:    ":1000",
		Handler: handler,
	}
	server.ListenAndServe()
}

func handle_test(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Fprint(w, "Dominus Iesus Christus")
}

func write_request_path(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.Path)
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

func more_specific_middleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("do a different specific middleware thing")
		h(w, r)
	}
}

func guardian_middleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("endpoint is protected")
		specific_middleware(h)(w, r)
	}
}

func prefix_middleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mux := http.NewServeMux()
		mux.HandleFunc("/test", handle_test)

		fmt.Println(r.URL.Path)
		next := http.StripPrefix("/prefix", mux)

		next.ServeHTTP(w, r)
	}
}
