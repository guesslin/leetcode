package main

import (
	"fmt"
)

func reverse(x int) int {
	var result int
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}
	for x != 0 {
		result = result*10 + x%10
		x = x / 10
		if result >= 2147483648 {
			return 0
		}
	}
	return sign * result
}

func main() {
	fmt.Println(reverse(-321))
}
