package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AskConfirmation(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt + " (y/n): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "y":
			return true
		case "n":
			return false
		default:
			fmt.Println("Invalid input, please type y or n.")
		}
	}
}
