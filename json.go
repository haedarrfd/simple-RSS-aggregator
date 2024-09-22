package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// respondWithError sends an HTTP response with an error message in JSON format
func respondWithError(w http.ResponseWriter, code int, msg string) {
	// Check if the internal server error
	if code > 499 {
		log.Println("Responding with 5XX error: ", msg)
	}

	// errorResponse struct to format the error message as a JSON format
	type errorResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errorResponse{Error: msg})
}

// respondWithJSON sends a JSON response to the client
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// Turn the payload to a JSON format
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response %v\n", payload)
		w.WriteHeader(500)
		return
	}

	// Add the content type to tell the client that the response is in JSON format,
	// send the HTTP status code,then write the JSON to response body
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
