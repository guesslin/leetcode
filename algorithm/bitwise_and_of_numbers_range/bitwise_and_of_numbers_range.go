package main

import (
	"fmt"
)

func rangeBitwiseAnd(m int, n int) int {
	if m == 0 {
		return 0
	}
	for n&(n-1) > m {
		n &= (n - 1)
	}
	return n & m
}

func main() {
	fmt.Println(5, 7, rangeBitwiseAnd(5, 7))
	fmt.Println(1, 13, rangeBitwiseAnd(1, 13))
	fmt.Println(13, 299, rangeBitwiseAnd(13, 299))
	fmt.Println(20000, 2147483647, rangeBitwiseAnd(20000, 2147483647))
}
