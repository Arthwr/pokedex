package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *Config) error
}

type LocationResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Config struct {
	Next     string
	Previous string
}

var commands map[string]cliCommand

func setupCommands() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Lists next 20 locations from Pokemon World",
			callback:    commandListNextLocations,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists previous 20 locations from Pokemon World",
			callback:    commandListPrevLocations,
		},
	}
}

func commandExit(c *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for _, cmd := range commands {
		fmt.Printf("  %-5s : %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandListNextLocations(c *Config) error {
	return fetchAndPrintLocations(c, c.Next)
}

func commandListPrevLocations(c *Config) error {
	return fetchAndPrintLocations(c, c.Previous)
}
