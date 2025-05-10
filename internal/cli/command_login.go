package cli

import (
	"context"
	"fmt"
)

func LoginHandler(s *State, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("username is required")
	}

	username := cmd.Arguments[0]

	_, err := s.DB.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("error loading user: %v", err)
	}

	s.Config.SetUser(username)

	fmt.Printf("User set to %s\n", username)
	return nil
}
