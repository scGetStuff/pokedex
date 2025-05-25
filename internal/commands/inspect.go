package commands

import (
	"fmt"
)

func comamndInspect(cmdArgs []string) error {
	if len(cmdArgs) < 1 {
		return fmt.Errorf("inspect command requires an argument")
	}

	poke, ok := pokemons[cmdArgs[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	// formating is inconsistant to match the bugs in the lesson
	fmt.Println("Name: ", poke.Name)
	fmt.Println("Height: ", poke.Height)
	fmt.Println("Weight: ", poke.Weight)
	fmt.Println("Stats: ")
	for _, stat := range poke.Stats {
		switch stat.Stat.Name {
		case "hp":
			fallthrough
		case "attack":
			fallthrough
		case "defense":
			fallthrough
		case "special-attack":
			fallthrough
		case "special-defense":
			fallthrough
		case "speed":
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		default:
		}
	}
	fmt.Println("Types: ")
	for _, t := range poke.Types {
		fmt.Println("  - ", t.Type.Name)
	}

	return nil
}
