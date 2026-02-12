package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	following, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	fmt.Println("you are following: ")
	for _, follow := range following {
		fmt.Printf("- %s\n", follow.FeedName)
	}

	return nil

}
