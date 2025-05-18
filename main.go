package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/scGetStuff/pokedex/internal/commands"
)

func main() {
	// fmt.Println("Hello, World!")

	const prompt = "Pokedex > "
	cmds := commands.GetCommandsMap()
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

			command, ok := cmds[first]
			if !ok {
				fmt.Println("Unknown command")
				continue
			}

			err := command.Callback()
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
