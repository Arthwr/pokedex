package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(c *config, params ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandHelp(c *config, params ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")

	commands := getCommands()
	nameWidth := findLongestCommandName(commands)

	for _, cmd := range commands {
		fmt.Printf(" - %-*s : %s\n", nameWidth, cmd.name, cmd.description)
	}

	return nil
}

func commandListNextLocations(c *config, params ...string) error {
	locations, err := c.pokeapiClient.FetchLocations(c.nextLocationURL)
	if err != nil {
		return err
	}

	c.nextLocationURL = locations.Next
	c.previousLocationURL = locations.Previous

	for _, loc := range locations.Results {
		fmt.Printf(" - %s\n", loc.Name)
	}

	return nil
}

func commandListPrevLocations(c *config, params ...string) error {
	if c.previousLocationURL == nil {
		return errors.New("you're on the first page")
	}

	locations, err := c.pokeapiClient.FetchLocations(c.previousLocationURL)
	if err != nil {
		return err
	}

	c.nextLocationURL = locations.Next
	c.previousLocationURL = locations.Previous

	for _, loc := range locations.Results {
		fmt.Printf(" - %s\n", loc.Name)
	}

	return nil
}

func commandExplore(c *config, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("usage: explore <location-name>")
	}

	locationName := params[0]
	encounters, err := c.pokeapiClient.FetchEncountersFromLocation(locationName)
	if err != nil {
		return err
	}

	if len(encounters.PokemonEncounters) == 0 {
		fmt.Printf("No Pokemon found in location %q.\n", locationName)
		return nil
	}

	for _, er := range encounters.PokemonEncounters {
		fmt.Printf("  - %s\n", er.Pokemon.Name)
	}

	return nil
}

func commandCatch(c *config, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("usage: catch <pokemon-name>")
	}

	pokemonName := params[0]
	pokemon, err := c.pokeapiClient.FetchPokemon(pokemonName)
	if err != nil {
		return err
	}

	return nil
}
