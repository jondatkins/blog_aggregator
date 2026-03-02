package main

import "context"

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
	s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	return nil
}
