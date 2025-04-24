package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("Closing the Pokedex... Goodbye!")
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for key, value := range commandRegistry {
		fmt.Printf("%s: %v\n", key, value.description)
	}
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commandRegistry = map[string]cliCommand{}
