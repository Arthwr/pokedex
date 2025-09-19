package main

import (
	"time"

	"github.com/arthwr/pokedex/internal/pokeapi"
)

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(10*time.Second, 20*time.Second),
	}

	runREPL(cfg)
}
