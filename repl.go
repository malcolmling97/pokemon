package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startRepl() {
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
				err := command.callback()
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
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		// "map": {
		// 	name:        "map",
		// 	description: "Displays next 20 Pokemon map",
		// 	callback:    MapCommand,
		// },
		// "mapb": {
		// 	name:        "mapb",
		// 	description: "Displays previous 20 Pokemon map",
		// 	callback:    MapBackCommand,
		// },
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

// // pokeapi/types.go
// type LocationAreasResponse struct {
// 	Count    int     `json:"count"`
// 	Next     *string `json:"next"`
// 	Previous *string `json:"previous"`
// 	Results  []struct {
// 		Name string `json:"name"`
// 		URL  string `json:"url"`
// 	} `json:"results"`
// }

// // config/config.go
// type Config struct {
// 	NextLocationAreaURL     *string
// 	PreviousLocationAreaURL *string
// }

// func MapCommand(cfg *Config) error {
// 	// URL to request
// 	url := "https://pokeapi.co/api/v2/location-area"
// 	if cfg.NextLocationAreaURL != nil {
// 		url = *cfg.NextLocationAreaURL
// 	}

// 	// Make HTTP request
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	// Parse response
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return err
// 	}

// 	var locAreaResp LocationAreasResponse
// 	err = json.Unmarshal(body, &locAreaResp)
// 	if err != nil {
// 		return err
// 	}

// 	// Update config with new URLs
// 	cfg.NextLocationAreaURL = locAreaResp.Next
// 	cfg.PreviousLocationAreaURL = locAreaResp.Previous

// 	// Display the location areas
// 	for _, area := range locAreaResp.Results {
// 		fmt.Println(area.Name)
// 	}

// 	return nil
// }

// // commands/mapb.go (simplified example)
// func MapBackCommand(cfg *Config) error {
// 	if cfg.PreviousLocationAreaURL == nil {
// 		fmt.Println("You're on the first page")
// 		return nil
// 	}

// 	// Make HTTP request to previous URL
// 	resp, err := http.Get(*cfg.PreviousLocationAreaURL)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	// Parse response
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return err
// 	}

// 	var locAreaResp LocationAreasResponse
// 	err = json.Unmarshal(body, &locAreaResp)
// 	if err != nil {
// 		return err
// 	}

// 	// Update config with new URLs
// 	cfg.NextLocationAreaURL = locAreaResp.Next
// 	cfg.PreviousLocationAreaURL = locAreaResp.Previous

// 	// Display the location areas
// 	for _, area := range locAreaResp.Results {
// 		fmt.Println(area.Name)
// 	}

// 	return nil
// }
