package main

import (
	"fmt"
	"time"
)

// Update the agg command to now take a single argument: time_between_reqs.
// time_between_reqs is a duration string, like 1s, 1m, 1h, etc. I used the time.ParseDuration function to parse it into a time.Duration value.
// It should print a message like Collecting feeds every 1m0s when it starts.
// Use a time.Ticker to run your scrapeFeeds function once every time_between_reqs. I used a for loop to ensure
// that it runs immediately (I don't like waiting) and then every time the ticker ticks:
// TechCrunch: https://techcrunch.com/feed/
// Hacker News: https://news.ycombinator.com/rss
// Boot.dev Blog: https://www.boot.dev/blog/index.xml
func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: agg <time_duration> e.g. agg 5s, agg 1m ...")
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}
	ticker := time.NewTicker(timeBetweenRequests)
	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)

	for ; ; <-ticker.C {
		fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)
		scrapeFeeds(s)
	}
	// feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	// if err != nil {
	// 	return fmt.Errorf("Couldn't get feed: %w", err)
	// }
	// fmt.Printf("Feed: %+v\n", feed)
	// return nil
}
