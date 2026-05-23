package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"gibhub.com/jzross/rssagg/internal/database"
)

//it will be on a go routine. it'll search for feeds table to get the oldests feeds that were updated and
//fetch its datato the database

func ScraperTriggerRoutine(q *database.Queries, ReqInterval time.Duration, concurrency int32) {

	var wg sync.WaitGroup

	ticker := time.NewTicker(ReqInterval)
	for ; ; <-ticker.C {

		feeds, err := q.GetLastFetched(context.Background(), concurrency)
		if err != nil {
			log.Printf("Error in getting last fetched: %v", err)
			continue
		}
		for _, feed := range feeds {
			wg.Add(1)
			go scraper(q, feed, &wg)

		}
		wg.Wait() // blocks waiting untill all wg.Adds are Done
	}
}

func scraper(q *database.Queries, feed database.Feed, wg *sync.WaitGroup) {
	defer wg.Done()

	// update the last update col for rss;
	feedUpdate, err := q.UpdateLastFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Some error while updating the feed's last updated date: %v", err)
	}
	fmt.Printf("Feed updated: %v", feedUpdate)

	// get data with the url from feed struct.
	rssfeed, err := getRSSFeed(feed.Url)
	if err != nil {
		log.Printf("Error in getting RSS: %v", err)
	}

	fmt.Print(rssfeed)
}
