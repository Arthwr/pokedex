package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/arthwr/pokedex/internal/pokeapi"
	"github.com/arthwr/pokedex/internal/pokestorage"
)

type config struct {
	pokeapiClient       pokeapi.Client
	pokemonStorage      *pokestorage.Storage
	nextLocationURL     *string
	previousLocationURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config, params ...string) error
}

func runREPL(cfg *config) {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}

		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}

		commandName := text[0]
		cmd, exists := commands[commandName]
		if !exists {
			fmt.Println("Unknown command. Type 'help' for available commands.")
			continue
		}

		params := text[1:]
		if err := cmd.callback(cfg, params...); err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "Explore location area by name and list all Pokemon found there. Usage: explore <location-name>",
			callback:    commandExplore,
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
		"catch": {
			name:        "catch",
			description: "Attempt to catch requested Pokemon",
			callback:    commandCatch,
		},
		"list": {
			name:        "list",
			description: "List all currently caught Pokemon",
			callback:    commandList,
		},
	}
}
