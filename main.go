package main

import (
	"fmt"

	"github.com/jondatkins/blog_aggregator/internal/config"
)

type state struct {
	conf *config.Config
}

type command struct {
	name string
	args []string
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}
	err = cfg.SetUser("jon")
	if err != nil {
		panic(err)
	}
	configContents, err := config.Read()
	if err != nil {
		panic(err)
	}
	fmt.Println("", configContents)
}
