package main

import (
	"fmt"
)

func getSum(a int, b int) int {
	if b == 0 {
		return a
	}
	carry := (a & b) << 1
	return getSum(a^b, carry)
}

func main() {
	a, b := 1, 2
	fmt.Println(a, "+", b, "=", getSum(a, b))
}
