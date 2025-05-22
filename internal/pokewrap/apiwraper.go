package pokewrap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/scGetStuff/pokedex/internal/pokecache"
)

const (
	baseURL  = "https://pokeapi.co/api/v2"
	pageSize = 20
)

var cache = pokecache.NewCache(5 * time.Second)

// mapPage: 0 based multiplier for offset
func GetLocationAreaJSON(mapPage int) (LocationArea, error) {
	// https://pokeapi.co/api/v2/location-area
	// https://pokeapi.co/api/v2/location-area?limit=20&offset=0

	url := fmt.Sprintf("%s/location-area?limit=%v&offset=%v", baseURL, pageSize, pageSize*mapPage)

	var data []byte
	data, hit := cache.Get(url)
	if !hit {
		resp, err := http.Get(url)
		if err != nil {
			return LocationArea{}, err
		}
		defer resp.Body.Close()

		if resp.StatusCode > 299 {
			return LocationArea{}, fmt.Errorf("response failed with status code: %d", resp.StatusCode)
		}
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationArea{}, err
		}
		cache.Add(url, data)
	}

	area := LocationArea{}
	if err := json.Unmarshal(data, &area); err != nil {
		return LocationArea{}, err
	}

	fmt.Println(url)

	// fmt.Print(area)
	// fmt.Print(pretty(string(data)))

	return area, nil
}

func Pretty(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}

	return prettyJSON.String(), nil
}
