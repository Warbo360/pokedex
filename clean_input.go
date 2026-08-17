package main

import "strings"

func cleanInput(text string) []string {

	loweredText := strings.ToLower(text)
	splitText := strings.Fields(loweredText)
	return splitText

}
