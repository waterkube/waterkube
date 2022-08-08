package cmd

import (
	"fmt"

	"github.com/petaki/support-go/cli"
)

// DiverExplore function.
func DiverExplore(group *cli.Group, command *cli.Command, arguments []string) int {
	fmt.Println()
	fmt.Println("  🤿 Swimming to " + cli.Green("coordinate") + "...")
	fmt.Println()

	fmt.Println("  ✅ Excavation has " + cli.Green("begun"))
	fmt.Println()

	return cli.Success
}

// DiverHire function.
func DiverHire(group *cli.Group, command *cli.Command, arguments []string) int {
	fmt.Println()
	fmt.Println("  💼 Looking for an " + cli.Green("applicant") + "...")
	fmt.Println()

	fmt.Println("  ✅ Diver is " + cli.Green("ready"))
	fmt.Println()

	return cli.Success
}
