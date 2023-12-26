package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("error unmashalling data %v", err)
		w.WriteHeader(500)
		return
	}
	w.Write(data)
	w.WriteHeader(code)

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondBody := struct {
		Error string `json:"error"`
	}{
		Error: message,
	}
	respondWithJson(w, code, respondBody)
}
