package main

import (
	"bufio"
	"fmt"
	"os"
	"pokemon/internal/pokeapi"
	"strings"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !reader.Scan() {
			break
		}

		arrText := cleanInput(reader.Text())

		if len(arrText) > 0 {
			firstword := arrText[0]

			command, exists := getCommands()[firstword]

			if exists {
				err := command.callback(cfg)
				if err != nil {
					fmt.Printf("Error: %v", err)
				}

			} else {
				fmt.Println("Unknown Command")
			}
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Displays next 20 Pokemon map",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 Pokemon map",
			callback:    commandMapb,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
