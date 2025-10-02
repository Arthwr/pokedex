package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/arthwr/pokedex/internal/pokecache"
)

type Client struct {
	httpClient   http.Client
	pokeCache    pokecache.Cache
	pokemonStore map[string]PokemonResponse
}

func (c *Client) doRequest(url string, target any) error {
	if val, ok := c.pokeCache.Get(url); ok {
		return json.Unmarshal(val, target)
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected HTTP status: %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	c.pokeCache.Add(url, data)
	return json.Unmarshal(data, target)
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
