package pokeapi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/arthwr/pokedex/internal/pokecache"
)

type Client struct {
	httpClient   http.Client
	pokeCache    pokecache.Cache
	pokemonStore map[string]PokemonResponse
}

func (c *Client) StorePokemon(pr PokemonResponse) {
	c.pokemonStore[pr.Name] = pr
}

func (c *Client) ListPokemons() {
	if len(c.pokemonStore) == 0 {
		fmt.Println("No Pokemon have been caught.")
		return
	}

	fmt.Println("List of stored Pokemon:")
	for p := range c.pokemonStore {
		fmt.Printf("  %s\n", p)
	}
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeCache:    *pokecache.NewCache(cacheInterval),
		pokemonStore: make(map[string]PokemonResponse),
	}
}
