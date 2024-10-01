package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/haedarrfd/simple-rss-aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	// JSON decoder that read request body
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	// Parsed that request body
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	// Add new feed into feeds table
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		Url:       params.URL,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
	})
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't add new feed")
		return
	}

	responseWithJSON(w, http.StatusOK, databaseFeedToFeed(feed))
}

// Retrieve all feeds that are already stored in the database
func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't get feeds")
		return
	}

	responseWithJSON(w, http.StatusOK, databaseFeedsToFeeds(feeds))
}
