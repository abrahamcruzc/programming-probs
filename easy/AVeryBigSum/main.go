package main

func main() {

}

func aVeryBigSum(ar []int64) int64 {
	var long int64

	for _, number := range ar {
		long += number
	}

	return long
}