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
						Name:       "create",
						Usage:      "Create a map",
						HandleFunc: cmd.MapCreate,
					},
					{
						Name:       "delete",
						Usage:      "Delete the map",
						HandleFunc: cmd.MapDelete,
					},
				},
			},
			{
				Name:  "artifact",
				Usage: "Artifact commands",
				Commands: []*cli.Command{
					{
						Name:       "combine",
						Usage:      "Combine the artifacts",
						Arguments:  []string{"artifactName1", "artifactName2"},
						HandleFunc: cmd.ArtifactCombine,
					},
					{
						Name:       "donate",
						Usage:      "Donate the artifact",
						Arguments:  []string{"artifactName"},
						HandleFunc: cmd.ArtifactDonate,
					},
					{
						Name:       "sell",
						Usage:      "Sell the artifact",
						Arguments:  []string{"artifactName"},
						HandleFunc: cmd.ArtifactSell,
					},
				},
			},
			{
				Name:  "diver",
				Usage: "Diver commands",
				Commands: []*cli.Command{
					{
						Name:       "explore",
						Usage:      "Explore the coordinate",
						Arguments:  []string{"gridName"},
						HandleFunc: cmd.DiverExplore,
					},
					{
						Name:       "hire",
						Usage:      "Hire a diver",
						HandleFunc: cmd.DiverHire,
					},
				},
			},
			{
				Name:  "submarine",
				Usage: "Submarine commands",
				Commands: []*cli.Command{
					{
						Name:       "explore",
						Usage:      "Explore the coordinate",
						Arguments:  []string{"gridName"},
						HandleFunc: cmd.SubmarineExplore,
					},
					{
						Name:       "buy",
						Usage:      "Buy a submarine",
						HandleFunc: cmd.SubmarineBuy,
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
