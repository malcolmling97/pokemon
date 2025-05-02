package pokeapi

type RespLocationPokemons struct {
	LocationName      string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type RespPokemonInfo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Base_EXP int    `json:"base_experience"`
	Height   int    `json:"height"`
	Weight   int    `json:"weight"`
}

// pokedex can reference the id
type Pokedex struct {
	Entries map[int]RespPokemonInfo
}
