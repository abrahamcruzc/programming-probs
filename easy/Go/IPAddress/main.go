package main

import (
	"fmt"
	"strings"
)

func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}

func main() {
	ip := "192.168.0.1"
	defanged := defangIPaddr(ip)
	fmt.Println(defanged) // Output: 192[.]168[.]0[.]1
}