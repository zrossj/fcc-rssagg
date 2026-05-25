package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"gibhub.com/jzross/rssagg/internal/database"
	"github.com/google/uuid"
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
	_, err := q.UpdateLastFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Some error while updating the feed's last updated date: %v", err)
	}

	// get data with the url from feed struct.
	rssfeed, err := getRSSFeed(feed.Url)
	if err != nil {
		log.Printf("Error in getting RSS: %v", err)
	}

	for _, item := range rssfeed.Channel.Item {
		descrip := sql.NullString{}
		if item.Description != "" {
			descrip.String = item.Description
			descrip.Valid = true
		}
		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Couldn't parse date %v to desired format. Error: %v", item.PubDate, err)
		}

		_, err = q.InsertPosts(context.Background(),
			database.InsertPostsParams{
				ID:          uuid.New(),
				CreatedAt:   time.Now().UTC(),
				UpdatedAt:   time.Now().UTC(),
				Title:       item.Title,
				Description: descrip,
				PublishedAt: pubDate, // need to cast
				Url:         item.Link,
				FeedID:      feed.ID,
			})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Printf("Error while trying to fetch post on db: %v", err)
		}
	}
	log.Printf("Feed %v collected. Found %v posts", feed.Name, len(rssfeed.Channel.Item))
}
