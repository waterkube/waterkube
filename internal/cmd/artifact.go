package cmd

import (
	"fmt"

	"github.com/petaki/support-go/cli"
)

// ArtifactCombine function.
func ArtifactCombine(group *cli.Group, command *cli.Command, arguments []string) int {
	fmt.Println()
	fmt.Println("  🏺 Finding the required " + cli.Green("artifacts") + "...")
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

	err = gameManager.MapLoad()
	if err != nil {
		return command.PrintError(err)
	}

	fmt.Println()
	fmt.Println("  📦 Sending the artifact to " + cli.Green("museum") + "...")
	fmt.Println()

	err = gameManager.ArtifactDonate(parsed[0])
	if err != nil {
		return command.PrintError(err)
	}

	fmt.Println("  🚢 New ship is " + cli.Green("on its way") + "...")
	fmt.Println()

	fmt.Println("  ✅ Artifact " + cli.Green("donated"))
	fmt.Println()

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

	err = gameManager.MapLoad()
	if err != nil {
		return command.PrintError(err)
	}

	fmt.Println()
	fmt.Println("  🤑 Payment " + cli.Green("in progress") + "...")
	fmt.Println()

	err = gameManager.ArtifactSell(parsed[0])
	if err != nil {
		return command.PrintError(err)
	}

	fmt.Println("  📦 Sending the artifact to " + cli.Green("customer") + "...")
	fmt.Println()

	fmt.Println("  ✅ Artifact " + cli.Green("sold"))
	fmt.Println()

	return cli.Success
}
