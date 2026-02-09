package main

import (
	"log"
	"os"

	"github.com/Kazimlyc/blog-aggregator/internal/config"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	programState := &state{
		cfg: &cfg,
	}

	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		log.Fatal("not enough arguments")
	}

	cmdName := args[1]
	cmdArgs := args[2:]
	cmd := command{name: cmdName, args: cmdArgs}

	if err := cmds.run(programState, cmd); err != nil {
		log.Fatal(err)
	}

}
