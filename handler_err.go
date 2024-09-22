package main

import "net/http"

// handlerErr sends a client-side error to an HTTP request
func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong")
}
