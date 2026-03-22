package cmd

import (
	"fmt"

	"github.com/petaki/support-go/cli"
	"github.com/waterkube/waterkube/internal/config"
	"github.com/waterkube/waterkube/internal/service"
)

// MapCreate function.
func MapCreate(group *cli.Group, command *cli.Command, arguments []string) int {
	appConfig, err := config.NewConfig(command, arguments)
	if err != nil {
		return command.PrintHelp(group)
	}

	redisPool := service.RedisPool(appConfig)
	defer redisPool.Close()

	gameManager := service.GameManager(appConfig, redisPool)

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
	appConfig, err := config.NewConfig(command, arguments)
	if err != nil {
		return command.PrintHelp(group)
	}

	redisPool := service.RedisPool(appConfig)
	defer redisPool.Close()

	gameManager := service.GameManager(appConfig, redisPool)

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
