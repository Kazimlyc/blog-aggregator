package main

import (
	"context"
	"fmt"

	"github.com/Kazimlyc/blog-aggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {

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
