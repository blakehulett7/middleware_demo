package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

func r_forwarding_test(w http.ResponseWriter, r *http.Request) {
	recorder := httptest.NewRecorder()
	fetch_bentkey(recorder, r)
	resp := recorder.Result()

	var SampleResponse struct {
		Token string `json:"token"`
	}
	json.NewDecoder(resp.Body).Decode(&SampleResponse)
	fmt.Println(SampleResponse)
}
