package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Getting the API data
func getPokeAPI() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area?offset=0&limit=20")
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var data PokeData
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		fmt.Println("Error decoding json:", err)
		return err
	}
	// fmt.Println(data)
	for _, v := range data.Results {
		fmt.Printf("The names of the locations: %s\n", v.Name)
	}
	return nil
}

// Commands struct
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// Commands callbacks
func cmdHelp() error {
	fmt.Println("This is the list of commands you can use.")
	for cmd := range cliCmd {
		fmt.Println(cliCmd[cmd].name + ": " + cliCmd[cmd].description)
	}
	return fmt.Errorf("Error HELP")
}

func cmdExit() error {
	fmt.Println("Test EXIT")
	os.Exit(0)
	// I have to figure out how to make an optional return, as this one is not going to happen due to exiting the program
	return fmt.Errorf("Error EXIT")
}

func cmdMap() error {
	fmt.Println("Moving up MAP")
	return fmt.Errorf("No more locations on the map")
}

func cmdMapB() error {
	fmt.Println("Moving back up MAP")
	return fmt.Errorf("No more locations on the map")
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

	fmt.Println("TEST") // NOTE: get rid of this

	// NOTE: testing API
	getPokeAPI()

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
			cliCmd["help"].callback()
		case cliCmd["exit"].name:
			cliCmd["exit"].callback()
		case cliCmd["map"].name:
			cliCmd["map"].callback()
		case cliCmd["mapb"].name:
			cliCmd["mapb"].callback()
		default:
			fmt.Println("If you need help with the commands, please type in \"help\" and hit ENTER")
		}
	}
}
