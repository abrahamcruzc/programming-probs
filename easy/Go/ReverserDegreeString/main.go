package main

import "fmt"

func main() {
	fmt.Println(reverseDegree("abc"))
}

func reverseDegree(s string) int {
	var degree int
	for i, char := range s {
		degree += int('z'-char+1) * (i + 1)
	}

	return degree
}
