package cli

import (
	"context"
	"fmt"
	"time"

	"gator/internal/database"

	"github.com/google/uuid"
)

func AddFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.Arguments) < 2 {
		return fmt.Errorf("name and url are required")
	}

	feedname := cmd.Arguments[0]
	feedurl := cmd.Arguments[1]

	ctx := context.Background()

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedname,
		Url:       feedurl,
		UserID:    user.ID,
	}

	feed, err := s.DB.CreateFeed(ctx, params)
	if err != nil {
		return fmt.Errorf("error creating feed: %v", err)
	}
	fmt.Printf("Registered new feed: %+v\n", feed)

	followCmd := Command{
		Arguments: []string{feed.Url},
	}
	err = Follow(s, followCmd, user)
	if err != nil {
		return fmt.Errorf("error following the feed: %w", err)
	}

	return nil
}
