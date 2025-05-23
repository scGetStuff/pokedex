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

func getBytes(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return []byte{}, fmt.Errorf("response failed with status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

// mapPage: 0 based multiplier for offset
func GetLocationAreaJSON(mapPage int) (LocationArea, error) {
	// https://pokeapi.co/api/v2/location-area
	// https://pokeapi.co/api/v2/location-area?limit=20&offset=0

	url := fmt.Sprintf("%s/location-area?limit=%v&offset=%v", baseURL, pageSize, pageSize*mapPage)

	var data []byte
	var hit bool
	data, hit = cache.Get(url)
	if !hit {
		// this is important
		// `data, err :=` would create a block variable and fuck everything up
		var err error
		data, err = getBytes(url)
		if err != nil {
			return LocationArea{}, err
		}

		cache.Add(url, data)
	}
	fmt.Println(string(data))

	area := LocationArea{}
	if err := json.Unmarshal(data, &area); err != nil {
		return LocationArea{}, err
	}

	// fmt.Println(url)

	// fmt.Print(area)
	// fmt.Print(PrettyJSON(string(data)))

	return area, nil
}

func PrettyJSON(str string) (string, error) {
	var bytesBuffer bytes.Buffer
	if err := json.Indent(&bytesBuffer, []byte(str), "", "    "); err != nil {
		return "", err
	}

	return bytesBuffer.String(), nil
}
