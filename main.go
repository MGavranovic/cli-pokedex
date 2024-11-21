package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/MGavranovic/cli-pokedex/pokeCache"
	"github.com/MGavranovic/cli-pokedex/pokeapi"
)

var (
	currentUrl string = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
)

// Commands struct
type cliCommand struct {
	name        string
	description string
	callback    func(c *Config, cache pokecache.Cache) error
}

// Commands callbacks
func cmdHelp(c *Config, cache pokecache.Cache) error {
	fmt.Println("This is the list of commands you can use.")
	for cmd := range cliCmd {
		fmt.Println("- " + cliCmd[cmd].name + ": " + cliCmd[cmd].description)
	}
	return fmt.Errorf("Error trying to use \"help\"")
}

func cmdExit(c *Config, cache pokecache.Cache) error {
	fmt.Println("Test EXIT")
	os.Exit(0)
	// I have to figure out how to make an optional return, as this one is not going to happen due to exiting the program
	return fmt.Errorf("Error EXIT")
}

// DEBUG:
// if on the last page or on the first, return err instead of trying to get the url first

func cmdMap(c *Config, cache pokecache.Cache) error {
	url := currentUrl
	if c.Next != nil {
		url = *c.Next
	}

	var data *pokeapi.PokeData

	if cachedData, found := cache.Get(url); found {
		if err := json.Unmarshal(cachedData, &data); err != nil {
			return err
		}
	} else {
		var err error
		data, err = pokeapi.GetPokeLocations(url)
		if err != nil {
			return err
		}

		rawData, err := json.Marshal(data)
		if err != nil {
			return err
		}
		cache.Add(url, rawData)
	}

	for _, v := range data.Results {
		fmt.Printf("The names of the locations: %s\n", v.Name)
	}

	c.Next = &data.Next
	c.Previous = &data.Previous

	return nil
}

func cmdMapB(c *Config, cache pokecache.Cache) error {
	if c.Previous == nil {
		return fmt.Errorf("You are already on the first page")
	}

	var data *pokeapi.PokeData

	if cachedData, found := cache.Get(*c.Previous); found {
		fmt.Println("Cache hit: ", *c.Previous)
		if err := json.Unmarshal(cachedData, &data); err != nil {
			return err
		}
	} else {
		fmt.Println("Cache miss, fetching from API:", *c.Previous)
		var err error
		data, err := pokeapi.GetPokeLocations(*c.Previous)
		if err != nil {
			return err
		}

		rawData, err := json.Marshal(data)
		if err != nil {
			return err
		}
		cache.Add(*c.Previous, rawData)
	}

	for _, v := range data.Results {
		fmt.Printf("The names of the locations: %s\n", v.Name)
	}

	c.Previous = &data.Previous
	c.Next = &data.Next

	return nil
}

// map cli commands
var cliCmd map[string]cliCommand

// init cli commands
func init() {
	cliCmd = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "About the app. List and description of all commands.",
			callback:    cmdHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the application.",
			callback:    cmdExit,
		},
		"map": {
			name:        "map",
			description: "Move up on the map.",
			callback:    cmdMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Move back on the map",
			callback:    cmdMapB,
		},
	}
}

func main() {
	fmt.Println("Hello there!")
	fmt.Println("Welcome to CLI Pokedex.")
	fmt.Println("A CLI tool where you will be able to see and learn all about your favourite pokemons.")
	cache := pokecache.NewCache(5 * time.Second)
	config := &Config{}

	// user input
	var userInput string

	// continuous loop
	for {
		fmt.Print("pokedex > ")

		// ask for user input
		fmt.Scan(&userInput)

		// respond to user input
		switch userInput {
		case cliCmd["help"].name:
			cliCmd["help"].callback(config, cache)
		case cliCmd["exit"].name:
			cliCmd["exit"].callback(config, cache)
		case cliCmd["map"].name:
			if err := cliCmd["map"].callback(config, cache); err != nil {
				fmt.Printf("An error occurred while trying to use \"%v\": %v\n", cliCmd["map"].name, err)
			}
		case cliCmd["mapb"].name:
			if err := cliCmd["mapb"].callback(config, cache); err != nil {
				fmt.Printf("An error occurred while trying to use \"%v\": %v\n", cliCmd["mapb"].name, err)
			}
		default:
			fmt.Println("If you need help with the commands, please type in \"help\" and hit ENTER")
		}
	}
}
