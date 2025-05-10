package cli

import (
	"context"
	"fmt"
)

func Users(s *State, _ Command) error {

	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error loading user: %v", err)
	}

	for _, user := range users {

		isCurrent := ""

		if s.Config.User == user.Name {
			isCurrent = "(current)"
		}
		fmt.Printf("* %v %v\n", user.Name, isCurrent)
	}

	return nil
}
