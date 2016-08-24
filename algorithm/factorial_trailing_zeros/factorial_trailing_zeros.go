package main

import (
	"fmt"
)

func trailingZeroes(n int) int {
	// count 5 in factor
	var five int
	for n/5 > 0 {
		five += n / 5
		n = n / 5
	}
	return five
}

func main() {
	fmt.Println(trailingZeroes(125))
}
