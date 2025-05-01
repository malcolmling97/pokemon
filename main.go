package main

import (
	"pokemon/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 10*time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)

}
