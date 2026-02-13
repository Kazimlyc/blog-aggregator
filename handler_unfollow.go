package main

import (
	"context"
	"fmt"

	"github.com/Kazimlyc/blog-aggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	feedurl := cmd.Args[0]
	feedId, err := s.db.GetFeedByURL(context.Background(), feedurl)
	if err != nil {
		return err
	}

	err = s.db.DeleteFollow(context.Background(), database.DeleteFollowParams{
		UserID: user.ID,
		FeedID: feedId.ID,
	})

	return err
}
