package main

// func handlerFollowMine(s *state, cmd command) error {
// 	if len(cmd.Args) != 1 {
// 		return fmt.Errorf("usage: %v <Feed URL>", cmd.Name)
// 	}
//
// 	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
// 	if err != nil {
// 		return fmt.Errorf("couldn't find user: %w", err)
// 	}
// 	feedURL := cmd.Args[0]
//
// 	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)
// 	if err != nil {
// 		return fmt.Errorf("couldn't find feed: %w", err)
// 	}
// 	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
// 		ID:        uuid.New(),
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 		UserID:    currentUser.ID,
// 		FeedID:    feed.ID,
// 	})
// 	if err != nil {
// 		return fmt.Errorf("couldn't create feed follow: %w", err)
// 	}
//
// 	fmt.Println("Feed Follow created")
// 	printFeedFollow(feedFollow)
// 	return nil
// }
//
// // Add a following command. It should print all the names of the feeds the current user is following.
// func handlerFollowing(s *state, cmd command) error {
// 	// if len(cmd.Args) != 1 {
// 	// 	return fmt.Errorf("usage: %v <user name>", cmd.Name)
// 	// }
//
// 	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
// 	if err != nil {
// 		return fmt.Errorf("Couldn't get user: %w", err)
// 	}
// 	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
// 	if err != nil {
// 		return fmt.Errorf("Couldn't get feed follows: %w", err)
// 	}
// 	for _, feedFollow := range feedFollows {
// 		fmt.Printf("- '%s'\n", feedFollow.FeedName)
// 	}
// 	return nil
// }
//
// func printFollowing(feed database.Feed) {
// 	fmt.Printf(" * ID:			%v\n", feed.ID)
// 	fmt.Printf(" * Name:			%v\n", feed.Name)
// }
//
// // print name of feed and current user
// func printFeedFollow(feedFollow database.CreateFeedFollowRow) {
// 	fmt.Printf(" * Name:			%v\n", feedFollow.FeedName)
// 	fmt.Printf(" * ID:			%v\n", feedFollow.UserName)
// }
