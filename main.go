package main

import (
	"time"

	"github.com/arthwr/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(10 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	runREPL(cfg)
}
