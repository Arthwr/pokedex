package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for _, cmd := range getCommands() {
		fmt.Printf("  %-5s : %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandListNextLocations(c *config) error {
	locations, err := c.pokeapiClient.FetchLocations(c.nextLocationURL)
	if err != nil {
		return err
	}

	c.nextLocationURL = locations.Next
	c.previousLocationURL = locations.Previous

	for _, loc := range locations.Results {
		fmt.Printf("  %s\n", loc.Name)
	}

	return nil
}

func commandListPrevLocations(c *config) error {
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
		fmt.Printf("  %s\n", loc.Name)
	}

	return nil
}
