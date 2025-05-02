package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, key *string) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationResp.Next
	cfg.prevLocationURL = locationResp.Prev

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config, key *string) error {
	if cfg.prevLocationURL == nil {
		return errors.New("you're on the first page\n")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationResp.Next
	cfg.prevLocationURL = locationResp.Prev

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
