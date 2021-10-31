package main

import "github.com/petaki/support-go/cli"

func main() {
	(&cli.App{
		Name:    "Waterkube",
		Version: "1.0.0",
	}).Execute()
}
