package main

import (
	"fmt"
	"net/http"

	"github.com/haedarrfd/simple-rss-aggregator/internal/auth"
	"github.com/haedarrfd/simple-rss-aggregator/internal/database"
)

// Type definition for an HTTP handler function that requires authentication
type authedHandler func(http.ResponseWriter, *http.Request, database.User)

// middlewareAuth ensures that only requests with a valid API key can access the wrapped handler
func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the API key from the request headers
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Couldn't find the api key: %v", err))
			return
		}

		// Find the user associated with the API key
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusNotFound, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
