package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/petaki/support-go/cli"
	"github.com/waterkube/waterkube/internal/cmd"
)

func main() {
	(&cli.App{
		Name:    "Waterkube",
		Version: "1.0.0",
		Groups: []*cli.Group{
			{
				Name:  "map",
				Usage: "Map commands",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "Create a map",
					},
					{
						Name:  "delete",
						Usage: "Delete the map",
					},
				},
			},
			{
				Name:  "artifact",
				Usage: "Artifact commands",
				Commands: []*cli.Command{
					{
						Name:  "combine",
						Usage: "Combine the artifacts",
					},
					{
						Name:  "donate",
						Usage: "Donate the artifact",
					},
					{
						Name:  "sell",
						Usage: "Sell the artifact",
					},
				},
			},
			{
				Name:  "diver",
				Usage: "Diver commands",
				Commands: []*cli.Command{
					{
						Name:  "explore",
						Usage: "Explore the coordinate",
					},
					{
						Name:  "hire",
						Usage: "Hire a diver",
					},
				},
			},
			{
				Name:  "submarine",
				Usage: "Submarine commands",
				Commands: []*cli.Command{
					{
						Name:  "explore",
						Usage: "Explore the coordinate",
					},
					{
						Name:  "buy",
						Usage: "Buy a submarine",
					},
				},
			},
			{
				Name:  "web",
				Usage: "Web commands",
				Commands: []*cli.Command{
					{
						Name:       "serve",
						Usage:      "Serve the app",
						HandleFunc: cmd.WebServe,
					},
				},
			},
		},
	}).Execute()
}
