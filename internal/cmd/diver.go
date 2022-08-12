package cmd

import (
	"fmt"

	"github.com/petaki/support-go/cli"
)

// DiverExplore function.
func DiverExplore(group *cli.Group, command *cli.Command, arguments []string) int {
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
	fmt.Println("  🤿 Swimming to " + cli.Green("coordinate") + "...")
	fmt.Println()

	err = gameManager.DiverExplore(parsed...)
	if err != nil {
		return command.PrintError(err)
	}

	fmt.Println("  ✅ Excavation has " + cli.Green("begun"))
	fmt.Println()

	return cli.Success
}

// DiverHire function.
func DiverHire(group *cli.Group, command *cli.Command, arguments []string) int {
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
	fmt.Println("  💼 Looking for an " + cli.Green("applicant") + "...")
	fmt.Println()

	err = gameManager.DiverHire()
	if err != nil {
		return command.PrintError(err)
	}

	fmt.Println("  ✅ Diver is " + cli.Green("ready"))
	fmt.Println()

	return cli.Success
}
