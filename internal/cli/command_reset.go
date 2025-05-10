package cli

import (
	"context"
	"fmt"
)

func Reset(s *State, _ Command) error {

	s.DB.DeleteUserTable(context.Background())
	fmt.Printf("Erased all users")
	return nil
}
