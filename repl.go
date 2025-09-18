package main

import (
	"bufio"
	"fmt"
	"os"
)

func runREPL(cfg *Config) {
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

		cmd, ok := commands[text[0]]
		if !ok {
			fmt.Println("Unknown command. Type 'help' for available commands.")
			continue
		}

		if err := cmd.callback(cfg); err != nil {
			fmt.Println("Error:", err)
		}
	}
}
