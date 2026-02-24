package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jondatkins/blog_aggregator/internal/config"
	"github.com/jondatkins/blog_aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.cmds[cmd.name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}
	st := &state{
		cfg: &cfg,
	}
	_ = st

	dbURL := cfg.DBURL
	db, err := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)
	st.db = dbQueries
	cmds := &commands{
		cmds: make(map[string]func(*state, command) error),
	}
	_ = cmds

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	commandName := getArgs()[0]
	argsSlice := getArgs()[1:]
	err = cmds.run(st, command{name: commandName, args: argsSlice})
	if err != nil {
		fmt.Println("err is ", err)
	}
}

func getArgs() []string {
	if len(os.Args) < 2 {
		fmt.Println("One argument required")
		os.Exit(1)
	}
	return os.Args[1:] // os.Args[0] is the program name
	// return os.Args
}
