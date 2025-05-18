package cli

import (
	"context"
	"fmt"
)

func Feeds(s *State, _ Command) error {

	feeds, err := s.DB.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error loading feeds: %v", err)
	}

	for _, feed := range feeds {

		user, err := s.DB.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("error loading user: %v", err)
		}

		fmt.Printf("* %v %v %v * \n", feed.Name, feed.Url, user.Name)

	}

	return nil
}
