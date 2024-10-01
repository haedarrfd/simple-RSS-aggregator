package main

import (
	"net/http"

	"github.com/haedarrfd/simple-rss-aggregator/internal/auth"
	"github.com/haedarrfd/simple-rss-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

// Ensures that only requests with a valid API key can access the wrapped handler
func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the API key from headers
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			responseWithError(w, http.StatusUnauthorized, "Couldn't find the api key")
			return
		}

		// Retrieve a user based on an API key
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			responseWithError(w, http.StatusNotFound, "Couldn't get user")
			return
		}

		handler(w, r, user)
	}
}
