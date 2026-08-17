package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		if len(cleanedInput) == 0 {
			continue
		}
		fmt.Printf("Your command was: %v\n", cleanedInput[0])
	}
}

func cleanInput(text string) []string {

	loweredText := strings.ToLower(text)
	splitText := strings.Fields(loweredText)
	return splitText

}
