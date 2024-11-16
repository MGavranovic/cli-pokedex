package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

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

var cliCmd map[string]cliCommand

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

	var userInput string

	for {
		fmt.Print("pokedex > ")

		fmt.Scan(&userInput)

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
