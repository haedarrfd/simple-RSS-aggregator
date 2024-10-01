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

// Fetching and scraping feeds using multiple concurrent goroutines at the specified time intervals
func startScraping(db *database.Queries, concurency int, timeBetweenRequest time.Duration) {
	log.Printf("Collecting feeds every %s on %v goroutines...", timeBetweenRequest, concurency)
	ticker := time.NewTicker(timeBetweenRequest)

	// Infinite loop that runs every time a new value come across the ticker channel
	for ; ; <-ticker.C {
		// Grab the next batch of feeds to fetch
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurency))
		if err != nil {
			log.Println("Couldn't get next feeds to fetch", err)
			continue
		}

		// Synchronization to wait for all of goroutines to finish executing
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			// Spawn new goroutine to scrape the feed concurrently
			go scrapeFeed(db, wg, feed)
		}

		// Wait all all the goroutines have finished
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	// Mark the feed as fetched in the database
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed %s fetched: %v", feed.Name, err)
		return
	}

	// Scrape the feed
	rssFeed, err := fetchFeed(feed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %s: %v", feed.Name, err)
		return
	}

	// Loop each individualy posts
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
