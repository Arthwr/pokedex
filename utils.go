package main

import (
	"math/rand"
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

func canCatchPokemon(baseExp int, maxBaseExp int, minChance, maxChance float64) bool {
	normalized := float64(baseExp) / float64(maxBaseExp)
	chance := maxChance - (maxChance-minChance)*normalized

	if chance > maxChance {
		chance = maxChance
	}

	return rand.Float64() < chance
}
