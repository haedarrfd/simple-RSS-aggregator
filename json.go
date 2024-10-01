package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Sends an HTTP response with an error message in JSON
func responseWithError(w http.ResponseWriter, code int, errMsg string) {
	// Internal server error
	if code > 499 {
		log.Println("Responding with 5XX error: ", errMsg)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	responseWithJSON(w, code, errorResponse{Error: errMsg})
}

// Sends an HTTP response with the payload in JSON
func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// Add the content type in JSON format
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	// Turn the payload to a JSON format
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON %v\n", err)
		w.WriteHeader(500)
		return
	}

	w.Write(data)
}
