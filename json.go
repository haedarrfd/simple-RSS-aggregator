package main

import (
	"encoding/json"
	"log"
	"net/http"
)

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
