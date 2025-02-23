package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error:", msg);
	}
	// below struct tells us that in struct we have Error field and `json:"error"` what will be the key in response
	type errResponse struct {
		Error string `json:"error"`
	}
	errorMsg := errResponse{
		Error: msg,
	}
	respondWithJSON(w, code, errorMsg)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload);
		w.WriteHeader(500)
		return;
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}