package main

import (
	"fmt"
	"os"

	"github.com/doddy-s/miren/cmd/miren/commands"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No command provided")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "upgoing", "up":
		commands.Upgoing()
	case "downgoing", "down":
		commands.Downgoing()
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
}
