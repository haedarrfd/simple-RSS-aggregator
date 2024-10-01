package main

import (
	"fmt"
	"net/http"

	"github.com/haedarrfd/simple-rss-aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	// Query to get posts for the user with a limit of 10 posts
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get posts: %v", err))
		return
	}

	respondWithJSON(w, 200, databasePostsToPosts(posts))
}
