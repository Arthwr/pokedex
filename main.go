package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		text := cleanInput(scanner.Text())
		fmt.Println("Your command was:", text[0])
	}
}

func cleanInput(text string) []string {
	formatted := strings.ToLower(text)
	words := strings.Fields(formatted)
	return words
}
