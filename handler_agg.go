package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {

	ctx := context.Background()

	fetch, err := fetchFeed(ctx, "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Printf("Feed: %+v\n", fetch)

	return nil

}
