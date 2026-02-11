package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {

		fmt.Printf("- %s\n", feed.Name)
		fmt.Printf(" - %s\n", feed.Url)
		user, err := s.db.GetUserById(context.Background(), feed.UserID.UUID)
		if err != nil {
			return err
		}
		fmt.Printf(" - %s\n", user)
	}

	return nil
}
