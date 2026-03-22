package cmd

import (
	"fmt"
	"strconv"

	"github.com/petaki/support-go/cli"
	"github.com/waterkube/waterkube/internal/config"
	"github.com/waterkube/waterkube/internal/game"
	"github.com/waterkube/waterkube/internal/service"
)

// SubmarineExplore function.
func SubmarineExplore(group *cli.Group, command *cli.Command, arguments []string) int {
	appConfig, err := config.NewConfig(command, arguments)
	if err != nil {
		return command.PrintHelp(group)
	}

	parsed := command.FlagSet().Args()
	redisPool := service.RedisPool(appConfig)
	defer redisPool.Close()

	gameManager := service.GameManager(appConfig, redisPool)

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
	appConfig, err := config.NewConfig(command, arguments)
	if err != nil {
		return command.PrintHelp(group)
	}

	redisPool := service.RedisPool(appConfig)
	defer redisPool.Close()

	gameManager := service.GameManager(appConfig, redisPool)

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
