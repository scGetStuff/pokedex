package commands

import (
	"fmt"
)

func commandPokedex(cmdArgs []string) error {

	fmt.Println("Your Pokedex:")
	for k := range pokemons {
		fmt.Println(" - ", k)
	}

	return nil
}
