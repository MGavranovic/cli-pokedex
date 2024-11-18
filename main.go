package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var (
	currentUrl string = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
)

// Could use pointers maybe the value doesnt update aaa
// could use bool params for next and previous (if true go next ...)
// Getting the API data
func getPokeAPI(url string, showLocations bool) (*PokeData, error) {
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

// Commands struct
type cliCommand struct {
	name        string
	description string
	callback    func(c *Config) error
}

// Commands callbacks
func cmdHelp(c *Config) error {
	fmt.Println("This is the list of commands you can use.")
	for cmd := range cliCmd {
		fmt.Println("- " + cliCmd[cmd].name + ": " + cliCmd[cmd].description)
	}
	return fmt.Errorf("Error trying to use \"help\"")
}

func cmdExit(c *Config) error {
	fmt.Println("Test EXIT")
	os.Exit(0)
	// I have to figure out how to make an optional return, as this one is not going to happen due to exiting the program
	return fmt.Errorf("Error EXIT")
}

func cmdMap(c *Config) error {
	url := currentUrl
	if c.Next != nil {
		url = *c.Next
	} else {
	}

	data, err := getPokeAPI(url, true)
	if err != nil {
		return err
	}

	c.Next = &data.Next
	c.Previous = &data.Previous

	return nil
}

func cmdMapB(c *Config) error {
	if c.Previous == nil {
		return fmt.Errorf("You are already on the first page")
	}
	data, err := getPokeAPI(*c.Previous, true)
	if err != nil {
		return err
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
	config := &Config{}

	fmt.Println("TEST") // NOTE: get rid of this

	// NOTE: testing API
	// getPokeAPI()

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
			cliCmd["help"].callback(config)
		case cliCmd["exit"].name:
			cliCmd["exit"].callback(config)
		case cliCmd["map"].name:
			if err := cliCmd["map"].callback(config); err != nil {
				fmt.Printf("An error occurred while trying to use \"%v\": %v\n", cliCmd["map"].name, err)
			}
		case cliCmd["mapb"].name:
			if err := cliCmd["mapb"].callback(config); err != nil {
				fmt.Printf("An error occurred while trying to use \"%v\": %v\n", cliCmd["mapb"].name, err)
			}
		default:
			fmt.Println("If you need help with the commands, please type in \"help\" and hit ENTER")
		}
	}
}
