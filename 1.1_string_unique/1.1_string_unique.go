package main

import (
	"fmt"
)

func main() {
	words := []string{"abcde", "test", "ぁあぃぎじ", "せぬがだぬ", ""}

	for _, word := range words {
		fmt.Printf("Is \"%s\" unique? %t\n", word, unique(word))
	}
}

func unique(input string) bool {
	seen := make(map[rune]bool)

	for _, c := range input {
		// If we've seen this character before, the string is not unique.
		if _, ok := seen[c]; ok {
			return false
		}

		seen[c] = true
	}

	return true
}
