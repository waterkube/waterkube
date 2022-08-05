package cmd

import (
	"fmt"

	"github.com/petaki/support-go/cli"
)

// SubmarineExplore function.
func SubmarineExplore(group *cli.Group, command *cli.Command, arguments []string) int {
	fmt.Println()
	fmt.Println("  📡 Navigating to " + cli.Green("coordinate") + "...")
	fmt.Println()

	fmt.Println("  ✅ Excavation has " + cli.Green("begun"))
	fmt.Println()

	return cli.Success
}

// SubmarineBuy function.
func SubmarineBuy(group *cli.Group, command *cli.Command, arguments []string) int {
	fmt.Println()
	fmt.Println("  📦 Ordering a " + cli.Green("submarine") + "...")
	fmt.Println()

	fmt.Println("  ✅ Submarine is " + cli.Green("ready"))
	fmt.Println()

	return cli.Success
}
