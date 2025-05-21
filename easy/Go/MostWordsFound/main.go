package main

import (
	"fmt"
	"strings"
)

func main() {
	sentences := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	fmt.Println(mostWordsFound(sentences))
}

func mostWordsFound(sentences []string) int {
	var maxWords int
	for _, sentence := range sentences {
		sentenceLen := len(strings.Split(sentence, " "))
		if sentenceLen > maxWords {
			maxWords = sentenceLen 
		}
	}
	return maxWords
}