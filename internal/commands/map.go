package commands

import (
	"fmt"

	"github.com/scGetStuff/pokedex/internal/pokewrap"
)

var mapPage = -1

func commandMap() error {
	area, err := pokewrap.GetLocationAreaJSON(mapPage + 1)
	if err != nil {
		return err
	}

	mapPage++

	for i := range area.Results {
		fmt.Println(area.Results[i].Name)
	}

	return nil
}

func commandMapb() error {
	if mapPage < 1 {
		return nil
	}

	area, err := pokewrap.GetLocationAreaJSON(mapPage - 1)
	if err != nil {
		return err
	}

	mapPage--

	for i := range area.Results {
		fmt.Println(area.Results[i].Name)
	}

	return nil
}
