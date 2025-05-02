package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(cfg *config, key *string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("Closing the Pokedex... Goodbye!")
}
