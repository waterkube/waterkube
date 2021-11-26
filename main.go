package main

import (
	"github.com/petaki/support-go/cli"
	"github.com/petaki/waterkube/internal/cmd"
)

func main() {
	(&cli.App{
		Name:    "Waterkube",
		Version: "1.0.0",
		Groups: []*cli.Group{
			{
				Name:  "artifact",
				Usage: "Artifact commands",
				Commands: []*cli.Command{
					{
						Name:  "combine",
						Usage: "Combine the artifacts",
					},
					{
						Name:  "sell",
						Usage: "Sell an artifact",
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
