package main

import (
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, key *string) error {

	pokemonResp, err := cfg.pokeapiClient.CatchPokemon(key)
	if err != nil {
		return err
	}

	// storing into pokedex is here, but we haven't set up a retrieval
	cfg.pokeapiClient.Pokedex.Entries[pokemonResp.ID] = pokemonResp

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonResp.Name)

	// start rolling here:
	rand.Seed(time.Now().UnixNano())

	baseExp := pokemonResp.Base_EXP

	const minExp = 40  // anything ≤ this is 100% capture
	const maxExp = 340 // anything ≥ this is 1% capture

	var chance int

	if baseExp <= minExp {
		chance = 100
	} else if baseExp >= maxExp {
		chance = 1
	} else {
		// Linearly scale down from 100% to 1%
		// Formula: chance = 100 - scaled_decrease
		scaled := float64(baseExp-minExp) / float64(maxExp-minExp)
		chance = 100 - int(scaled*99) // 99 because 100 → 1 range
	}

	roll := rand.Intn(100) + 1 // 1–100

	if roll <= chance {
		fmt.Printf("%s was caught!\n", pokemonResp.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemonResp.Name)
	}

	return nil
}
