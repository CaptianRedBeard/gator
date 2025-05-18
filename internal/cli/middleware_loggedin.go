package cli

import (
	"context"
	"fmt"
	"gator/internal/database"
)

func MiddlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {

		ctx := context.Background()
		user, err := s.DB.GetUser(ctx, s.Config.User) // Assuming s.Config.User contains the user ID or username
		if err != nil {
			return fmt.Errorf("error loading user: %w", err)
		}
		if user.Name == "" {
			return fmt.Errorf("user is not logged in")
		}

		return handler(s, cmd, user)
	}
}
