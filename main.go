package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	words := strings.Split(text, " ")

	out := make([]string, 0, len(words))

	// fmt.Println("")
	for _, word := range words {
		// fmt.Printf("'%s', ", word)
		word = strings.TrimSpace(word)
		if word != "" {
			out = append(out, word)
		}
	}

	return out
}
