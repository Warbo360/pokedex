package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

var cliCommands map[string]cliCommand

func init() {

	/* Map of available CLI commands for the program */

	cliCommands = map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exits Pokedex REPL",
			callback: exit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: help,
		},
	}
}

func startRepl() {

	/* Takes in input from the user and cleans it */

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		command := input[0]

	/* Matches command input to callback for that command from cliCommands */

		switch command {
		case "exit":
			cliCommands["exit"].callback()
		case "help":
			cliCommands["help"].callback()
		default:
			fmt.Println("Unknown Command")
		}
	}
}

func cleanInput(text string) []string {

	loweredText := strings.ToLower(text)
	splitText := strings.Fields(loweredText)
	return splitText

}

func exit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func help() error {
	fmt.Printf("Welcome to Pokedex!\n")
	fmt.Printf("Usage:\n")
	for name, cmd := range cliCommands {
		fmt.Printf("  - %v: %v\n", name, cmd.description)
	}
	return nil
}
