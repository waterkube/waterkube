package cmd

import (
	"fmt"

	"github.com/petaki/support-go/cli"
	"github.com/waterkube/waterkube/internal/config"
	"github.com/waterkube/waterkube/internal/service"
)

// ArtifactCombine function.
func ArtifactCombine(group *cli.Group, command *cli.Command, arguments []string) int {
	appConfig, err := config.NewConfig(command, arguments)
	if err != nil {
		return command.PrintHelp(group)
	}

	parsed := command.FlagSet().Args()
	redisPool := service.RedisPool(appConfig)
	defer redisPool.Close()

	gameManager := service.GameManager(appConfig, redisPool)

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
	appConfig, err := config.NewConfig(command, arguments)
	if err != nil {
		return command.PrintHelp(group)
	}

	parsed := command.FlagSet().Args()
	redisPool := service.RedisPool(appConfig)
	defer redisPool.Close()

	gameManager := service.GameManager(appConfig, redisPool)

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
	appConfig, err := config.NewConfig(command, arguments)
	if err != nil {
		return command.PrintHelp(group)
	}

	parsed := command.FlagSet().Args()
	redisPool := service.RedisPool(appConfig)
	defer redisPool.Close()

	gameManager := service.GameManager(appConfig, redisPool)

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
