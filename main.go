package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jondatkins/blog_aggregator/internal/config"
	"github.com/jondatkins/blog_aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := &commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("register", handlerRegister)
	cmds.register("login", handlerLogin)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerListFeeds)
	cmds.register("follow", handlerFollow)
	cmds.register("following", handlerListFeedFollows)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}
	commandName := os.Args[1]
	commandArgs := os.Args[2:]
	err = cmds.run(programState, command{Name: commandName, Args: commandArgs})
	if err != nil {
		log.Fatal(err)
	}
}

func getArgs() []string {
	if len(os.Args) < 2 {
		fmt.Println("One argument required")
		os.Exit(1)
	}
	return os.Args[1:] // os.Args[0] is the program name
}
