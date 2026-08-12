package clean_input
import "strings"

func cleanInput(text string) []string {

	split_text := strings.Split(strings.ToLower(text), " ")
	return split_text

}
