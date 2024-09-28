package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/haedarrfd/simple-rss-aggregator/internal/database"
)

// handlerCreateUser is a method that has access to apiConfig struct (database) to create a data
func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	// To capture the expected name from request body
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	// Create a JSON decoder that read request body
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	// Parsed that request body, if there's an error return client error
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// Add new user into users table in the database, if there's an error return client error
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		Url:       params.URL,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't add new user: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}
