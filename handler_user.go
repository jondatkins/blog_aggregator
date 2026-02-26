package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jondatkins/blog_aggregator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		// fmt.Println("username required for register")
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}
	name := cmd.Args[0]
	// user, err := s.db.GetUser(context.Background(), name)
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Println("User created")
	printUser(user)
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Printf("username switched to %s\n", name)
	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	currentUser := s.cfg.CurrentUserName
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Couldn't get users: %w", err)
	}
	for _, u := range users {
		if currentUser == u.Name {
			fmt.Printf("* '%s (current)'\n", u.Name)
		} else {
			fmt.Printf("* '%s'\n", u.Name)
		}
	}
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:			%v\n", user.ID)
	fmt.Printf(" * Name:			%v\n", user.Name)
}
