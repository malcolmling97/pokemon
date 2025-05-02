package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(location *string) (RespLocationPokemons, error) {

	url := baseURL + "/location-area/" + *location
	// this is caching method, to explore later once its up

	// if body, ok := c.cache.Get(url); ok {
	// 	locationsResp := RespShallowLocations{}
	// 	err := json.Unmarshal(body, &locationsResp)
	// 	if err != nil {
	// 		return RespShallowLocations{}, err
	// 	}

	// 	return locationsResp, nil
	// }

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationPokemons{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationPokemons{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationPokemons{}, err
	}

	pokemonResp := RespLocationPokemons{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return RespLocationPokemons{}, err
	}

	// c.cache.Add(url, data)

	return pokemonResp, nil

}

func (c *Client) CatchPokemon(pokemon *string) (RespPokemonInfo, error) {

	url := baseURL + "/pokemon/" + *pokemon

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemonInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemonInfo{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemonInfo{}, err
	}

	pokemonResp := RespPokemonInfo{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return RespPokemonInfo{}, err
	}

	return pokemonResp, nil

}
