package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func handlerAgg(s *state, cmd command) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.Name)
	}
	timeBetweenReqs := cmd.Args[0]

	fmt.Printf("Collecting feeds every %s\n", timeBetweenReqs)

	timeBetween, err := time.ParseDuration(timeBetweenReqs)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(timeBetween)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			log.Println("Error scraping feeds:", err)
		}
	}

}

func scrapeFeeds(s *state) error {

	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("\n--- fetching: %s (%s) ---\n", nextFeed.Name, nextFeed.Url)
	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return err
	}

	feedData, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	for _, item := range feedData.Channel.Item {
		fmt.Printf("%s\n", item.Title)

	}

	return nil
}
