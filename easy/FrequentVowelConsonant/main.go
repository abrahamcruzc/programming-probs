package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(maxFreqSum("successes"))
}

func maxFreqSum(s string) int {
	vowelStr := "aeiou"
	consonantStr := "bcdfghjklmnpqrstvwxyz"

	freq := make(map[rune]int)
	for _, char := range s {
		freq[char]++
	}

	maxVowel := 0
	maxConsonant := 0

	for char, count := range freq {
		if strings.ContainsRune(vowelStr, char) && count > maxVowel {
			maxVowel = count
		}
		if strings.ContainsRune(consonantStr, char) && count > maxConsonant {
			maxConsonant = count
		}
	}

	return maxVowel + maxConsonant
}