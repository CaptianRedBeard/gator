package cli

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"

	"gator/internal/database"
)

func RegisterHandler(s *State, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("username is required")
	}

	username := cmd.Arguments[0]
	_, err := s.DB.GetUser(context.Background(), username)
	if err == nil {
		return fmt.Errorf("username already exists")
	} else if err != sql.ErrNoRows {
		return fmt.Errorf("database error: %w", err)
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	user, err := s.DB.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}
	fmt.Printf("Registered new user: %+v\n", user)

	err = s.Config.SetUser(username)
	if err != nil {
		return fmt.Errorf("error setting user %v", err)
	}

	return nil
}
