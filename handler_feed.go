package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jondatkins/blog_aggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %v <Feed Name> <Feed URL>", cmd.Name)
	}

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}
	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]
	// user, err := s.db.GetUser(context.Background(), feedName)
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    currentUser.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Println("Feed created")
	printFeed(feed)
	return nil
}

func handlerGetFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Couldn't get users: %w", err)
	}
	for _, feed := range feeds {
		fmt.Printf("- '%s'\n", feed.Name)
		fmt.Printf("- '%s'\n", feed.Url)
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			fmt.Println("No User found for user id: ", feed.UserID)
		}
		fmt.Printf("- '%s'\n", user.Name)
	}
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf(" * ID:			%v\n", feed.ID)
	fmt.Printf(" * Name:			%v\n", feed.Name)
}
