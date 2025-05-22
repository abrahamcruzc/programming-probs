package main

import (
	"fmt"
	"strings"
)

func truncateSentence(s string, k int) string {
	strArr := strings.Split(s, " ")
	var builder strings.Builder

	for i := 0; i < k; i++ {
		builder.WriteString(strArr[i])
		if i < k-1 {
			builder.WriteString(" ")
		}
	}
	return builder.String()
}

func main() {
	fmt.Println(truncateSentence("Hello how are you Contestant", 4))
}