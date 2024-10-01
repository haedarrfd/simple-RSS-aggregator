package main

import (
	"net/http"

	"github.com/haedarrfd/simple-rss-aggregator/internal/database"
)

// Retrieve all posts for the user with a limit of 10 posts
func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't get posts")
		return
	}

	responseWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
