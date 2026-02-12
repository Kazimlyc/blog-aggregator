package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Kazimlyc/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFeed(s *state, cmd command) error {

	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	name := cmd.Args[0]
	url := cmd.Args[1]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      name,
			Url:       url,
			UserID: uuid.NullUUID{
				UUID:  user.ID,
				Valid: true,
			},
		},
	)
	if err != nil {
		return err
	}

	followFeed, err := s.db.CreateFeedFollow(context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			UserID:    user.ID,
			FeedID:    feed.ID,
		})
	if err != nil {
		return err
	}
	fmt.Printf("- %s, now following %s\n", followFeed.UserName, followFeed.FeedName)

	fmt.Printf("feed successfuly added: %s", feed.Name)

	return nil
}
