package commands

import (
	"fmt"
	"math/rand"

	"github.com/scGetStuff/pokedex/internal/pokewrap"
)

var pokemons = map[string]pokewrap.Pokemon{}

func commandCatch(cmdArgs []string) error {
	if len(cmdArgs) < 1 {
		return fmt.Errorf("catch command requires an argument")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", cmdArgs[0])

	vermin, err := pokewrap.GetPokemon(cmdArgs[0])
	if err != nil {
		return err
	}

	// TODO: no spec, I have no idea what the range is, so no idea how to use it
	base := vermin.BaseExperience
	_ = base
	num := rand.Intn(100)

	if num > 40 {
		pokemons[cmdArgs[0]] = vermin
		fmt.Println(cmdArgs[0], " was caught!")
	} else {
		fmt.Println(cmdArgs[0], " escaped!")
	}

	// for k := range pokemons {
	// 	fmt.Println(k)
	// }

	return nil
}
