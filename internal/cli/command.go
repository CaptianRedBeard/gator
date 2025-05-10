package cli

import (
	"fmt"
	"gator/internal/config"
	"gator/internal/database"
)

// State holds the application state, including a pointer to the configuration.
type State struct {
	DB     *database.Queries
	Config *config.Config
}

// Command represents a command with a name and a slice of arguments.
type Command struct {
	Name      string
	Arguments []string
}

// Commands holds the command handlers.
type Commands struct {
	handlers map[string]func(*State, Command) error
}

// NewCommands initializes a new Commands instance with an empty handlers map.
func NewCommands() *Commands {
	return &Commands{
		handlers: make(map[string]func(*State, Command) error),
	}
}

// Register adds a command handler to the Commands struct.
func (c *Commands) Register(name string, handler func(*State, Command) error) {
	c.handlers[name] = handler
}

// Run executes the command handler for the given command name.
func (c *Commands) Run(state *State, cmd Command) error {
	cmdFunc, ok := c.handlers[cmd.Name]
	if !ok {
		return fmt.Errorf("error running command: %s not found", cmd.Name)
	}

	return cmdFunc(state, cmd)
}
