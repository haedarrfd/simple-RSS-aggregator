package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

// The structure of an RSS feed in XML
type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}
type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

// urlToFeed fetch an RSS feed from the given URL
func urlToFeed(url string) (RSSFeed, error) {
	// Set a timeout of 10 seconds for the HTTP request
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	// Make an HTTP GET request to fetch the RSS feed from the provided URL
	response, err := httpClient.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}

	// Close response body after fetching
	defer response.Body.Close()

	// Read all response body
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return RSSFeed{}, err
	}

	rssFeed := RSSFeed{}

	// Unmarshal (parse) the XML data into rssFeed struct
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return RSSFeed{}, err
	}

	return rssFeed, nil
}
