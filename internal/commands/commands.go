package commands

import (
	"fmt"
	"os"
)

type CliCommand struct {
	Description string
	Callback    func() error
}

var commands map[string]CliCommand

func GetCommandsMap() map[string]CliCommand {
	// I want singleton behavior
	if commands == nil {
		// could not do this at declaration, circular dependency
		commands = map[string]CliCommand{
			"exit": {
				Description: "Exit the Pokedex",
				Callback:    commandExit,
			},
			"help": {
				Description: "Displays a help message",
				Callback:    commandHelp,
			},
		}

	}

	return commands
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for k, v := range commands {
		fmt.Printf("%s: %s\n", k, v.Description)
	}

	return nil
}
