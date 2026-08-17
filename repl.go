package main

import (
	"fmt"
	"bufio"
	"os"
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
