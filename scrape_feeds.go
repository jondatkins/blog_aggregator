package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jondatkins/blog_aggregator/internal/database"
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
	for _, foo := range feed.Channel.Item {
		fmt.Println(foo.Title)
		fmt.Println("=====================================")
	}
	return nil
}
