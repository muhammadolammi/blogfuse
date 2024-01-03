package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/muhammadolammi/blogfuse/internal/database"
)

func startScrapping(db *database.Queries, concurency int, timeBetweenRequest time.Duration) {
	log.Printf("Scrapping on %v goroutines in %v intervals", concurency, timeBetweenRequest)
	timeTicker := time.NewTicker(timeBetweenRequest)

	for ; ; <-timeTicker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurency))
		if err != nil {
			log.Println("there is an error getting next feeds, err:", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error marking feed %v as fetched, err: %v", feed.Name, err)
		return
	}

	rssFeed, err := urlToRssFeed(feed.Url)

	if err != nil {
		log.Printf("there is an error turning feed url %v to rssFeed, err: %v", feed.Name, err)
		return
	}
	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}
		if item.Description != "" {
			description = sql.NullString{
				Valid:  true,
				String: item.Description,
			}
		}
		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("there is an error parsing time on item %v, err: %v", item.Title, err)
		}
		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Description: description,
			Title:       item.Title,
			Url:         item.Link,
			PublishedAt: pubAt,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Printf("there is an error creating item post for item '%v', err: %v", item.Title, err)
		}
	}

	log.Printf("Feed %v collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))

}
