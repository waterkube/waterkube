package cmd

import (
	"fmt"

	"github.com/petaki/support-go/cli"
)

// MapCreate function.
func MapCreate(group *cli.Group, command *cli.Command, arguments []string) int {
	redisURL, redisKeyPrefix := createRedisFlags(command)

	_, err := command.Parse(arguments)
	if err != nil {
		return command.PrintHelp(group)
	}

	redisPool := newRedisPool(*redisURL)
	defer redisPool.Close()

	gameManager := newGameManager(*redisKeyPrefix, redisPool)

	fmt.Println()
	fmt.Println("  🐚 Creating a new " + cli.Green("map") + "...")
	fmt.Println()

	err = gameManager.MapCreate()
	if err != nil {
		return command.PrintError(err)
	}

	fmt.Println("  ✅ Map is " + cli.Green("ready"))
	fmt.Println()

	return cli.Success
}

// MapDelete function.
func MapDelete(group *cli.Group, command *cli.Command, arguments []string) int {
	redisURL, redisKeyPrefix := createRedisFlags(command)

	_, err := command.Parse(arguments)
	if err != nil {
		return command.PrintHelp(group)
	}

	redisPool := newRedisPool(*redisURL)
	defer redisPool.Close()

	gameManager := newGameManager(*redisKeyPrefix, redisPool)

	fmt.Println()
	fmt.Println("  ❌ Deleting the " + cli.Green("map") + "...")
	fmt.Println()

	err = gameManager.MapDelete()
	if err != nil {
		return command.PrintError(err)
	}

	fmt.Println("  ✅ Map is " + cli.Green("empty"))
	fmt.Println()

	return cli.Success
}
