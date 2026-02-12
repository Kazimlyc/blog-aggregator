package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Kazimlyc/blog-aggregator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func handlerUsers(s *state, cmd command) error {

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		fmt.Printf("* %s", user.Name)
		if user.Name == s.cfg.CurrentUserName {
			fmt.Print(" (current)")
		}
		fmt.Println()
	}

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
