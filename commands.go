package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Kazimlyc/blog-aggregator/internal/config"
	"github.com/Kazimlyc/blog-aggregator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("username is required")
	}

	username := cmd.Args[0]

	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return err
	}
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("username has been set to: %s\n", user.Name)

	return nil
}

func handlerRegister(s *state, cmd command) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	user, err := s.db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      name,
		},
	)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" { // unique_violation
				return fmt.Errorf("user already exists")
			}
		}
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return err
	}
	fmt.Printf("user registered: %s\n", user.Name)

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {

	c.handlers[name] = f

}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.handlers[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}
	return handler(s, cmd)
}
