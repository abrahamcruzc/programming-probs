package main

import (
	"fmt"
)

func countConsistentStrings(allowed string, words []string) int {
	allowedSet := make(map[rune]bool)
	for _, char := range allowed {
		allowedSet[char] = true
	}

	count := 0
	for _, word := range words {
		isConsistent := true

		for _, char := range word {
			if !allowedSet[char] {
				isConsistent = false
				break
			}
		}

		if isConsistent {
			count++
		}
	}

	return count
}

func main() {
	allowed := "ab"
	words := []string{"ad","bd","aaab","baa","badab"}
	fmt.Println(countConsistentStrings(allowed, words))
}