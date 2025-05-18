package cli

import (
	"context"
	"database/sql"
	"fmt"
	"gator/internal/database"
	"gator/internal/rss"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func Agg(s *State, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("time is required")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.Arguments[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %v", err)
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenReqs)
	ticker := time.NewTicker(timeBetweenReqs)
	defer ticker.Stop()

	if err := scrapeFeeds(s); err != nil {
		return fmt.Errorf("error running scrapeFeeds: %v", err)
	}

	for {
		select {
		case <-ticker.C:
			if err := scrapeFeeds(s); err != nil {
				fmt.Printf("error running scrapeFeeds: %v\n", err)
			}
		}
	}

}

func scrapeFeeds(s *State) error {

	ctx := context.Background()
	nextFeed, err := s.DB.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("faild fetching next feed: %v", err)
	}

	err = s.DB.MarkFeedFetched(ctx, nextFeed.ID)
	if err != nil {
		return fmt.Errorf("failed marking feed as fetched: %v", err)
	}

	feed, err := rss.FetchFeed(ctx, nextFeed.Url)
	if err != nil {
		return err
	}

	for _, item := range feed.Channel.Item {

		publishedAt, err := convertPublishedTime(item.PubDate)
		if err != nil {
			return err
		}

		params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: publishedAt,
			FeedID:      uuid.NullUUID{UUID: nextFeed.ID, Valid: true},
		}
		_, err = s.DB.CreatePost(ctx, params)
		if err != nil {
			if err, ok := err.(*pq.Error); ok {
				if err.Code == "23505" { // Unique violation
					continue
				}
			}
			return err
		}
	}

	return nil
}

func convertPublishedTime(timestamp string) (sql.NullTime, error) {
	layout := "Mon, 02 Jan 2006 15:04:05 -0700"

	parsedTime, err := time.Parse(layout, timestamp)
	if err != nil {
		return sql.NullTime{}, err
	}

	// Return sql.NullTime with the parsed time and valid set to true
	return sql.NullTime{
		Time:  parsedTime,
		Valid: true,
	}, nil
}
