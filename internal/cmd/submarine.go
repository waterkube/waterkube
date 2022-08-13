package cmd

import (
	"fmt"
	"strconv"

	"github.com/petaki/support-go/cli"
	"github.com/waterkube/waterkube/internal/game"
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

	fmt.Println()

	for _, gridName := range parsed {
		err = gameManager.MapLoad()
		if err != nil {
			return command.PrintError(err)
		}

		grid, err := gameManager.SubmarineExplore(gridName)
		if err != nil {
			return command.PrintError(err)
		}

		fmt.Println("  📡 Navigating to the " + cli.Green(grid.Name) + "...")
		fmt.Println()

		fmt.Println("  ✅ Excavation has " + cli.Green("begun"))
		fmt.Println()
	}

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

	err = gameManager.SubmarineBuy()
	if err != nil {
		if err == game.ErrNoMoney {
			fmt.Println("  🚫 You need " + cli.Yellow("$") + cli.Green(strconv.Itoa(game.SubmarinePrice)))

			return cli.Failure
		}

		return command.PrintError(err)
	}

	fmt.Println()
	fmt.Println("  📦 Ordering a " + cli.Green("submarine") + "...")
	fmt.Println()

	fmt.Println("  ✅ Submarine is " + cli.Green("ready"))
	fmt.Println()

	return cli.Success
}
