package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jondatkins/blog_aggregator/internal/database"
	"github.com/lib/pq"
)

// Write an aggregation function, I called mine scrapeFeeds. It should:
// Get the next feed to fetch from the DB.
// Mark it as fetched.
// Fetch the feed using the URL (we already wrote this function)
// Iterate over the items in the feed and print their titles to the console.
func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
		ID:            nextFeed.ID,
	})
	if err != nil {
		return err
	}
	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	for _, rssItem := range feed.Channel.Item {
		publishedAt, _ := time.Parse(time.RFC1123Z, rssItem.PubDate)
		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       rssItem.Title,
			Url:         rssItem.Link,
			Description: sql.NullString{String: rssItem.Description, Valid: rssItem.Description != ""},
			PublishedAt: sql.NullTime{Time: publishedAt, Valid: publishedAt.Year() > 1},
			FeedID:      nextFeed.ID,
		})
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok && pqErr.Code != "23505" {
				fmt.Printf("Error: %v\n", err)
			}
		}
	}
	return nil
}
