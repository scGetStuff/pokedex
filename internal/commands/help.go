package commands

import (
	"fmt"
)

func commandHelp(cmdArgs []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for k, v := range commands {
		fmt.Printf("%-*s: %s\n", helpNameWidth, k, v.Description)
	}

	return nil
}
