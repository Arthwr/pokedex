package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const pokemonLocationURL = "https://pokeapi.co/api/v2/location-area/"

func fetchAndPrintLocations(c *Config, url string) error {
	if url == "" {
		url = pokemonLocationURL
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected HTTP status: %s", res.Status)
	}

	var locations LocationResponse
	if err = json.NewDecoder(res.Body).Decode(&locations); err != nil {
		return err
	}

	if locations.Previous != nil {
		c.Previous = *locations.Previous
	}

	if locations.Next != nil {
		c.Next = *locations.Next
	}

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
