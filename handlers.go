// package main
//
// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"os"
// 	"time"
//
// 	"github.com/google/uuid"
// 	"github.com/jondatkins/blog_aggregator/internal/database"
// )
//
// func handlerLogin(s *state, cmd command) error {
// 	if len(cmd.args) < 1 {
// 		fmt.Println("username required for login")
// 		os.Exit(1)
// 	}
// 	_, err := s.db.GetUser(context.Background(), cmd.args[0])
// 	if err == sql.ErrNoRows {
// 		fmt.Println("No User with this name", cmd.args[0])
// 		os.Exit(1)
// 	} else if err != nil {
// 		fmt.Println("Error getting user: ", cmd.args[0], err)
// 	} else {
// 		fmt.Println("User exists: ", cmd.args[0])
// 		// os.Exit(1)
// 	}
// 	err = s.cfg.SetUser(cmd.args[0])
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("username set to %s\n", cmd.args[0])
// 	return nil
// }
//
// func handlerRegister(s *state, cmd command) error {
// 	if len(cmd.args) < 1 {
// 		fmt.Println("username required for register")
// 		os.Exit(1)
// 	}
// 	_, err := s.db.GetUser(context.Background(), cmd.args[0])
// 	if err == sql.ErrNoRows {
// 		s.db.CreateUser(context.Background(), database.CreateUserParams{
// 			ID:        uuid.New(),
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 			Name:      cmd.args[0],
// 		})
// 	} else if err != nil {
// 		fmt.Println("Error getting user: ", cmd.args[0], err)
// 	} else {
// 		fmt.Println("User exists: ", cmd.args[0])
// 		os.Exit(1)
// 	}
// 	err = s.cfg.SetUser(cmd.args[0])
// 	if err != nil {
// 		return err
// 	}
//
// 	fmt.Printf("username set to %s\n", cmd.args[0])
// 	return nil
// }
