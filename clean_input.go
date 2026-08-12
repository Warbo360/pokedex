package main

import "strings"

func cleanInput(text string) []string {

	trimmedText := strings.TrimSpace(text)
	loweredText := strings.ToLower(trimmedText)
	splitText := strings.Split(loweredText, " ")
	return splitText

}
