package main

import (
	"strconv"
	"strings"
	"fmt"
)

func convertDateToBinary(date string) string {
	dateArr := strings.Split(date, "-")
	year := dateArr[0]
	month := dateArr[1]
	day := dateArr[2]

	yearInt, _ := strconv.Atoi(year)
	monthInt, _ := strconv.Atoi(month)
	dayInt, _ := strconv.Atoi(day)

	yearBin := strconv.FormatInt(int64(yearInt), 2)
	monthBin := strconv.FormatInt(int64(monthInt), 2)
	dayBin := strconv.FormatInt(int64(dayInt), 2)

	return yearBin + "-" + monthBin + "-" + dayBin
}

func main() {
	date := "2080-02-29"
    result := convertDateToBinary(date)
    fmt.Printf("Input: %s\n", date)
    fmt.Printf("Output: %s\n", result)
    // Output: 100000100000-10-11101
}