package cli

import (
	"context"
	"fmt"
	"gator/internal/database"
)

func Following(s *State, _ Command, user database.User) error {

	following, err := s.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error loading : follows%w", err)
	}

	for _, feed := range following {
		fmt.Printf("* %v\n", feed.FeedName)
	}

	return nil
}
