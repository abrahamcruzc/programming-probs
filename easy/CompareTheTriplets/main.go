package main

import "fmt"

func main() {

	fmt.Println(compareTriplets([]int32{5,6,7}, []int32{3,6,10}))
}

func compareTriplets(a []int32, b []int32) []int32 {
	var alice int32
	var bob int32

	if len(a) != len(b) {
		return []int32{}
	}

	for i := range a {
		if a[i] > b[i] {
			alice += 1
		} else if b[i] > a[i] {
			bob += 1
		}
	}

	return []int32{alice, bob}
}