package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, key *string) error {

	if key == nil {
		return errors.New("exploring requires a location\n")
	}
	pokemonResp, err := cfg.pokeapiClient.ExploreLocation(key)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %v...\n", pokemonResp.LocationName)
	fmt.Println("Found Pokemon:")

	for _, encounter := range pokemonResp.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
