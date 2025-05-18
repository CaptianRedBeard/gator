package cli

import (
	"context"
	"fmt"

	"time"

	"gator/internal/database"

	"github.com/google/uuid"
)

func Follow(s *State, cmd Command, user database.User) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("url is required")
	}

	url := cmd.Arguments[0]

	ctx := context.Background()
	feed, err := s.DB.GetFeedByURL(ctx, url)
	if err != nil {
		return fmt.Errorf("error loading feed: %w", err)
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err = s.DB.CreateFeedFollow(ctx, params)
	if err != nil {
		return fmt.Errorf("error creating feed follow: %v", err)
	}
	fmt.Printf("Feed %+v followed by %+v\n", feed.Name, user.Name)

	return nil
}
