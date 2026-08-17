package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for x := 0; ; x++ {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
		cleanedInput := cleanInput(input)
		fmt.Printf("Your command was: %v\n", cleanedInput[0])
		continue
	}

}
