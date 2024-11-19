package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	currentUrl string = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
)

func GetPokeLocations(url string, showLocations bool) (*PokeData, error) {
	res, err := http.Get(url)
	if err != nil {
		return &PokeData{}, err
	}

	defer res.Body.Close()

	var data PokeData
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		fmt.Println("Error decoding json:", err)
		return &PokeData{}, err
	}

	if showLocations {
		for _, v := range data.Results {
			fmt.Printf("The names of the locations: %s\n", v.Name)
		}
	}
	return &data, nil
}
