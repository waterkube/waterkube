package cmd

import (
	"fmt"

	"github.com/petaki/support-go/cli"
)

// ArtifactCombine function.
func ArtifactCombine(group *cli.Group, command *cli.Command, arguments []string) int {
	redisURL, redisKeyPrefix := createRedisFlags(command)

	parsed, err := command.Parse(arguments)
	if err != nil {
		return command.PrintHelp(group)
	}

	redisPool := newRedisPool(*redisURL)
	defer redisPool.Close()

	gameManager := newGameManager(*redisKeyPrefix, redisPool)

	err = gameManager.MapLoad()
	if err != nil {
		return command.PrintError(err)
	}

	grid, err := gameManager.ArtifactCombine(parsed[0], parsed[1])
	if err != nil {
		return command.PrintError(err)
	}

	fmt.Println()
	fmt.Println("  🏺 Crafting the " + cli.Green(grid.Artifact) + "...")
	fmt.Println()

	fmt.Println("  ✅ Artifacts " + cli.Green("combined"))
	fmt.Println()

	return cli.Success
}

// ArtifactDonate function.
func ArtifactDonate(group *cli.Group, command *cli.Command, arguments []string) int {
	redisURL, redisKeyPrefix := createRedisFlags(command)

	parsed, err := command.Parse(arguments)
	if err != nil {
		return command.PrintHelp(group)
	}

	redisPool := newRedisPool(*redisURL)
	defer redisPool.Close()

	gameManager := newGameManager(*redisKeyPrefix, redisPool)

	fmt.Println()

	for _, artifactName := range parsed {
		err = gameManager.MapLoad()
		if err != nil {
			return command.PrintError(err)
		}

		grid, err := gameManager.ArtifactDonate(artifactName)
		if err != nil {
			return command.PrintError(err)
		}

		fmt.Println("  🚢 Museum has received the " + cli.Green(grid.Artifact) + "...")
		fmt.Println()

		fmt.Println("  ✅ Artifact " + cli.Green("donated"))
		fmt.Println()
	}

	return cli.Success
}

// ArtifactSell function.
func ArtifactSell(group *cli.Group, command *cli.Command, arguments []string) int {
	redisURL, redisKeyPrefix := createRedisFlags(command)

	parsed, err := command.Parse(arguments)
	if err != nil {
		return command.PrintHelp(group)
	}

	redisPool := newRedisPool(*redisURL)
	defer redisPool.Close()

	gameManager := newGameManager(*redisKeyPrefix, redisPool)

	fmt.Println()

	for _, artifactName := range parsed {
		err = gameManager.MapLoad()
		if err != nil {
			return command.PrintError(err)
		}

		grid, err := gameManager.ArtifactSell(artifactName)
		if err != nil {
			return command.PrintError(err)
		}

		fmt.Println("  🤑 A collector is paying for the " + cli.Green(grid.Artifact) + "...")
		fmt.Println()

		fmt.Println("  ✅ Artifact " + cli.Green("sold"))
		fmt.Println()
	}

	return cli.Success
}
