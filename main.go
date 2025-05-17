package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

func main() {
	// fmt.Println("Hello, World!")

	initMap()

	const prompt = "Pokedex > "
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)
		if input.Scan() {
			stuff := cleanInput(input.Text())
			if len(stuff) < 1 {
				continue
			}
			first := stuff[0]
			// fmt.Printf("Your command was: %s\n", first)

			command, ok := commands[first]
			if !ok {
				fmt.Println("Unknown command")
				continue
			}

			err := command.callback()
			if err != nil {
				fmt.Print(err)
			}
		}
	}

}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	words := strings.Split(text, " ")

	out := make([]string, 0, len(words))

	for _, word := range words {
		word = strings.TrimSpace(word)
		if word != "" {
			out = append(out, word)
		}
	}

	return out
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
		fmt.Printf("%s: %s\n", k, v.description)
	}

	return nil
}

// could not do this at declaration, circular dependency
func initMap() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}
