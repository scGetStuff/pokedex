package commands

import (
	"fmt"
	"math/rand"

	"github.com/scGetStuff/pokedex/internal/pokewrap"
)

// TODO: all the package variables should be pulled into a User struct
// first pass single thread, not a requirement, but this sucks
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
		pokemons[vermin.Name] = vermin
		fmt.Println(vermin.Name, " was caught!")
	} else {
		fmt.Println(vermin.Name, " escaped!")
	}

	return nil
}
