package cli

import (
	"context"
	"fmt"
	"gator/internal/database"
)

func Unfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("url is required")
	}

	url := cmd.Arguments[0]

	ctx := context.Background()
	feed, err := s.DB.GetFeedByURL(ctx, url)
	if err != nil {
		return err
	}

	params := database.DeleteFeedFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	}

	err = s.DB.DeleteFeedFollow(ctx, params)
	if err != nil {
		return fmt.Errorf("error deleting feed follow: %v", err)
	}
	fmt.Printf("Feed %+v unfollowed by %+v\n", feed.Name, user.Name)

	return nil
}
