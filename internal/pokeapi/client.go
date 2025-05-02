package pokeapi

import (
	"net/http"
	"pokemon/internal/pokecache"
	"time"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
	Pokedex    *Pokedex
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	pokeCache := pokecache.NewCache(cacheInterval)
	pokedex := &Pokedex{
		Entries: make(map[int]RespPokemonInfo),
	}

	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache:   pokeCache,
		Pokedex: pokedex,
	}
}
