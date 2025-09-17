package main

import (
	"strings"
)

func cleanInput(text string) []string {
	formatted := strings.ToLower(text)
	words := strings.Fields(formatted)
	return words
}
