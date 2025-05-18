package cli

import (
	"context"
	"fmt"
	"gator/internal/database"
	"strconv"
)

func Browse(s *State, cmd Command, user database.User) error {
	limit := 2
	if len(cmd.Arguments) > 0 {
		num, err := strconv.Atoi(cmd.Arguments[0])
		if err != nil {
			return fmt.Errorf("error converting arg to int: %v", cmd.Arguments[0])
		}
		limit = num
	}

	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	ctx := context.Background()
	posts, err := s.DB.GetPostsForUser(ctx, params)
	if err != nil {
		return nil
	}

	for _, post := range posts {
		fmt.Printf("ID: %v\nTitle: %v\nURL: %v\nDescription: %v\nPublished At: %v\n\n", post.ID, post.Title, post.Url, post.Description, post.PublishedAt)
	}

	return nil
}
