package main

import (
	"fmt"
	"math"
)

func main() {
	arr := [][]int32 {
		{1,2,3},
		{4,5,6},
		{9,8,9},
	}

	fmt.Println(diagonalDifference(arr))
}

func diagonalDifference(arr [][]int32) int32 {
	var left, right int32
	n := len(arr)

	for i := range arr {
		left += arr[i][i]
		right += arr[i][n-1-i]
	}

	return int32(math.Abs(float64(left - right)))
}