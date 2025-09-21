package main

import (
	"strings"
)

func cleanInput(text string) []string {
	formatted := strings.ToLower(text)
	words := strings.Fields(formatted)
	return words
}

func findLongestCommandName(commands map[string]cliCommand) int {
	maxLen := 0
	for _, cmd := range commands {
		if len(cmd.name) > maxLen {
			maxLen = len(cmd.name)
		}
	}
	return maxLen
}
