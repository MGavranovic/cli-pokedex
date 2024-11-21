package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	currentUrl string = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
)

func GetPokeLocations(url string) (*PokeData, error) {
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

	return &data, nil
}
