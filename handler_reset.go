package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		// fmt.Printf("Couldn't delete users: %w", err)
		// os.Exit(1)
		return fmt.Errorf("Couldn't delete users: %w", err)
	}
	fmt.Println("Database has been reset")
	return nil
}
