package main

import "net/http"

// handlerRead sends a response to an HTTP request
func handlerRead(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
