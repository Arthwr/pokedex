package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	formatted := strings.ToLower(text)
	words := strings.Fields(formatted)
	return words
}
