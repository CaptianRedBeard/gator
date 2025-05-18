package main

import (
	"database/sql"
	"fmt"
	"gator/internal/cli"
	"gator/internal/config"
	"gator/internal/database"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error with the config file:", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Println("Error with the database", err)
		os.Exit(1)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println("Error: could not connect to database:", err)
		os.Exit(1)
	}

	dbQueries := database.New(db)

	state := &cli.State{
		DB:     dbQueries,
		Config: &cfg,
	}

	// Initialize commands and register handlers
	commands := cli.NewCommands()
	commands.Register("login", cli.LoginHandler)
	commands.Register("register", cli.RegisterHandler)
	commands.Register("reset", cli.Reset)
	commands.Register("users", cli.Users)
	commands.Register("agg", cli.Agg)
	commands.Register("addfeed", cli.MiddlewareLoggedIn(cli.AddFeed))
	commands.Register("feeds", cli.Feeds)
	commands.Register("follow", cli.MiddlewareLoggedIn(cli.Follow))
	commands.Register("following", cli.MiddlewareLoggedIn(cli.Following))
	commands.Register("unfollow", cli.MiddlewareLoggedIn(cli.Unfollow))
	commands.Register("browse", cli.MiddlewareLoggedIn(cli.Browse))

	// Check if there are enough arguments
	if len(os.Args) < 2 {
		fmt.Println("Error: Not enough arguments provided")
		os.Exit(1)
	}

	// Create a command from the command-line arguments
	cmd := cli.Command{
		Name:      os.Args[1],
		Arguments: os.Args[2:],
	}

	// Run the command
	if err := commands.Run(state, cmd); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
