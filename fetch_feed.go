package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

// Write a func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) function. It should fetch a feed from the given URL, and, assuming that nothing goes wrong, return a filled-out RSSFeed struct. Here are some useful docs (be sure to check the Overviews for examples if the entry lacks any):
// http.NewRequestWithContext
// http.Client.Do
// I set the User-Agent header to gator in the request with request.Header.Set. This is a common practice to identify your program to the server.
// io.ReadAll
// xml.Unmarshal (works the same as json.Unmarshal)
// Use the html.UnescapeString function to decode escaped HTML entities (like &ldquo;). You'll need to run the Title and Description fields (of both the entire channel as well as the items) through this function.
// Add an agg command. Later this will be our long-running aggregator service. For now, we'll just use it to fetch a single feed and ensure our parsing works. It should fetch the feed found at https://www.wagslane.dev/index.xml and print the entire struct to the console.
func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var feed RSSFeed
	if err := xml.Unmarshal(data, &feed); err != nil {
		return nil, err
	}
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for _, item := range feed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}
	return &feed, nil
}

// Add an agg command. Later this will be our long-running aggregator
// service. For now, we'll just use it to fetch a single feed and ensure
// our parsing works. It should fetch the feed found at https://www.wagslane.dev/index.xml and print the entire struct to the console.
func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("Couldn't get feed: %w", err)
	}
	fmt.Println(feed)
	return nil
}
