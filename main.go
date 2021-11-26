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
