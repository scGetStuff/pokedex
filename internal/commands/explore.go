package commands

import (
	"fmt"

	"github.com/scGetStuff/pokedex/internal/pokewrap"
)

func commandExplore(cmdArgs []string) error {
	if len(cmdArgs) < 1 {
		return fmt.Errorf("explore command requires an argument")
	}

	encounter, err := pokewrap.GetLocationAreaEncounter(cmdArgs[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s\n", "TODO")
	fmt.Println("Found Pokemon:")

	for i := range encounter.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.PokemonEncounters[i].Pokemon.Name)
	}

	return nil
}
