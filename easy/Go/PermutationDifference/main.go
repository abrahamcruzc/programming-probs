package main

import (
	"fmt"
)

func main() {
	fmt.Println(findPermutationDifference("abc", "bac"))
}

func findPermutationDifference(s string, t string) int {
	pos := make(map[rune]int)
    for i, char := range t {
        pos[char] = i
    }

    var output int
    for i, char := range s {
        j := pos[char]
        diff := i - j
        if diff < 0 {
            diff = -diff
        }
        output += diff
    }

	return output
}