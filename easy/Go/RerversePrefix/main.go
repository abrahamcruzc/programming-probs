package main

import (
	"fmt"
	"strings"
)

func reversePrefix(word string, ch byte) string {
	occurrenceIndex := strings.IndexByte(word, ch)
	if occurrenceIndex == -1 {
		return word
	}
	wordBytes := []byte(word)
	for i := 0; i <= occurrenceIndex/2; i++ {
		wordBytes[i], wordBytes[occurrenceIndex-i] = wordBytes[occurrenceIndex-i], wordBytes[i]
	}

	return string(wordBytes)
}

func main() {
	fmt.Println(reversePrefix("abcdefd", 'd'))
	fmt.Println(reversePrefix("zxyxxe", 'z'))
	fmt.Println(reversePrefix("abcd", 'z'))
}
