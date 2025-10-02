package main

import (
	"time"

	"github.com/arthwr/pokedex/internal/pokeapi"
	"github.com/arthwr/pokedex/internal/pokestorage"
)

func main() {
	cfg := &config{
		pokeapiClient:  pokeapi.NewClient(10*time.Second, 20*time.Second),
		pokemonStorage: pokestorage.NewStorage(),
	}

	runREPL(cfg)
}
