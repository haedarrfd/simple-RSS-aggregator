package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/haedarrfd/simple-rss-aggregator/internal/database"
)

// Costum name formatting JSON
type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
	APIKey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

// databaseUserToUser populated with the corresponding fields from users table
func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdateAt:  dbUser.UpdatedAt,
		APIKey:    dbUser.ApiKey,
	}
}

// databaseFeedToFeed populated with the corresponding fields from feeds table
func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		Name:      dbFeed.Name,
		URL:       dbFeed.Url,
		CreatedAt: dbFeed.CreatedAt,
		UpdateAt:  dbFeed.UpdatedAt,
		UserID:    dbFeed.UserID,
	}
}

// databaseFeedsToFeeds converts into a slice of feeds
func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}

	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}

	return feeds
}

// databaseFeedFolToFeedFol populated with the corresponding fields from feed_follows table
func databaseFeedFolToFeedFol(dbFeedFol database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFol.ID,
		CreatedAt: dbFeedFol.CreatedAt,
		UpdateAt:  dbFeedFol.UpdatedAt,
		UserID:    dbFeedFol.UserID,
		FeedID:    dbFeedFol.FeedID,
	}
}
