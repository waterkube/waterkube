package cmd

import (
	"fmt"

	"github.com/petaki/support-go/cli"
)

// SubmarineExplore function.
func SubmarineExplore(group *cli.Group, command *cli.Command, arguments []string) int {
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
	fmt.Println("  📡 Navigating to " + cli.Green("coordinate") + "...")
	fmt.Println()

	err = gameManager.SubmarineExplore(parsed...)
	if err != nil {
		return command.PrintError(err)
	}

	fmt.Println("  ✅ Excavation has " + cli.Green("begun"))
	fmt.Println()

	return cli.Success
}

// SubmarineBuy function.
func SubmarineBuy(group *cli.Group, command *cli.Command, arguments []string) int {
	redisURL, redisKeyPrefix := createRedisFlags(command)

	_, err := command.Parse(arguments)
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
	fmt.Println("  📦 Ordering a " + cli.Green("submarine") + "...")
	fmt.Println()

	err = gameManager.SubmarineBuy()
	if err != nil {
		return command.PrintError(err)
	}

	fmt.Println("  ✅ Submarine is " + cli.Green("ready"))
	fmt.Println()

	return cli.Success
}
