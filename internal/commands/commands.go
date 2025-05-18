package commands

import (
	"fmt"
	"os"

	"github.com/scGetStuff/pokedex/internal/pokewrap"
)

// TODO: this shoud be a seperate file
type CliCommand struct {
	Description string
	Callback    func() error
}

// TOOD: I hate this idea
// add config param to callback function breaks everything
// so far, it is only used by the map/mapb commands; the same functionality
// not doing it untill it makes sense
type Config struct {
	NextURL     string
	PreviousURL string
}

var commands map[string]CliCommand

// to allign command names in the help function
var helpNameWidth = 0

func GetCommandsMap() map[string]CliCommand {
	// I want singleton behavior
	if commands == nil {
		// could not do this at declaration, circular dependency in Help
		commands = map[string]CliCommand{
			"exit": {
				Description: "Exit the Pokedex",
				Callback:    commandExit,
			},
			"help": {
				Description: "Displays a help message",
				Callback:    commandHelp,
			},
			"map": {
				Description: "Displays 20 Pokemon world locations at a time",
				Callback:    commandMap,
			},
			"mapb": {
				Description: "Displays 20 Pokemon world locations at a time",
				Callback:    commandMapb,
			},
		}

		for k := range commands {
			if helpNameWidth < len(k) {
				helpNameWidth = len(k)
			}
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
		fmt.Printf("%-*s: %s\n", helpNameWidth, k, v.Description)
	}

	return nil
}

var mapPage = -1

func commandMap() error {
	mapPage++

	area, err := pokewrap.GetMapJSON(mapPage)
	if err != nil {
		return err
	}

	for i := range area.Results {
		fmt.Println(area.Results[i].Name)
	}

	return nil
}

func commandMapb() error {
	if mapPage < 1 {
		return nil
	}

	mapPage--

	area, err := pokewrap.GetMapJSON(mapPage)
	if err != nil {
		return err
	}

	for i := range area.Results {
		fmt.Println(area.Results[i].Name)
	}

	return nil
}
