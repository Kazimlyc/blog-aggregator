package main

import (
	"fmt"

	"github.com/Kazimlyc/blog-aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {

	if len(cmd.args) == 0 {
		return fmt.Errorf("username is required")
	}

	username := cmd.args[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("username has been set to: %s\n", username)

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {

	c.handlers[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.handlers[cmd.name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}
	return handler(s, cmd)
}
