package main

import (
	"fmt"
	"strings"
)

func findWordsContaining(words []string, x byte) []int {
	var indices []int

	for i, word := range words {
		if strings.ContainsRune(word, rune(x)) {
			indices = append(indices, i)
		}
	}

	return indices
}

func main() {
	words := []string{"leet", "code"}
	x := byte('e')
	fmt.Println(findWordsContaining(words, x))
}