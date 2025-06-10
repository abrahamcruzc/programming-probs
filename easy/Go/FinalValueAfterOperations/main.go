package main

import (
	"fmt"
)

func finalValueAfterOperations(operations []string) int {
	var x int

	for _, operation := range operations {
		switch operation {
		case "++X", "X++":
			x++
		case "--X", "X--":
			x--
		}
	}
	return x
}

func main() {
	operations := []string{"++X","++X","X++"}
	fmt.Println(finalValueAfterOperations(operations))
}