package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/haedarrfd/simple-rss-aggregator/internal/database"
)

// startScraping fetching and scraping feeds using multiple concurrent goroutines at the specified time intervals
func startScraping(db *database.Queries, concurency int, timeBetweenRequest time.Duration) {
	log.Printf("Collecting feeds every %s on %v goroutines...", timeBetweenRequest, concurency)
	// Make a request on this interval
	ticker := time.NewTicker(timeBetweenRequest)

	// Infinite loop that runs every time a new value come across the ticker channel
	for ; ; <-ticker.C {
		// Grab the next batch of feeds to fetch
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurency))
		if err != nil {
			log.Println("Couldn't get next feeds to fetch", err)
			continue
		}

		// WaitGroup to synchronize the completion of all goroutines
		waitGrp := &sync.WaitGroup{}
		for _, feed := range feeds {
			// Add 1 to WaitGroup for every feed
			waitGrp.Add(1)

			// spawn new goroutine to scrape the feed concurrently
			go scrapeFeed(db, waitGrp, feed)
		}

		// Wait for all the spawned goroutines to complete before continue the loop
		waitGrp.Wait()
	}
}

// scrapeFeed scrapes a single feed from the provided URL
func scrapeFeed(db *database.Queries, waitGrp *sync.WaitGroup, feed database.Feed) {
	defer waitGrp.Done()

	// Mark the feed as fetched in the database
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed %s fetched: %v", feed.Name, err)
		return
	}

	// Scrape the feed
	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %s: %v", feed.Name, err)
		return
	}

	// Log each individualy post
	for _, item := range rssFeed.Channel.Item {
		// Handle description potential null values
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		// Parse published_at
		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Couldn't parse date %v with error %v", item.PubDate, err)
			continue
		}

		// Add new post into posts table
		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			Title:       item.Title,
			Description: description,
			PublishedAt: pubAt,
			Url:         item.Link,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			FeedID:      feed.ID,
		})
		if err != nil {
			// Skip the current post as it already exists
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println("Failed to create post:", err)
		}
	}

	log.Printf("Feed %s collected %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
