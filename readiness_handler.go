package main

import "net/http"

func readinessErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, "Internal Server Error")
}

func readinessSuccess(w http.ResponseWriter, r *http.Request) {
	respondBody := struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	}
	respondWithJson(w, 200, respondBody)
}
