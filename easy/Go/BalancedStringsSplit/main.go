package main

import "fmt"

func main() {
	fmt.Println(balancedStringSplit("LLLLRRRR"))
}

func balancedStringSplit(s string) int {
	balanceIndicator, pair := 0, 0
	for _, char := range s {
		if char == 'L' {
			balanceIndicator -= 1
		} else if char == 'R' {
			balanceIndicator += 1
		}
		if balanceIndicator == 0 {
			pair++
		}
	}
	return pair
}