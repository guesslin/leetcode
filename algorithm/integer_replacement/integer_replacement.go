package main

import (
	"fmt"
)

func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return 1 + integerReplacement(n/2)
	}
	if n == 2147483647 { // in case of 32 signed int overflow problem
		return min(1+integerReplacement(n-1), 2+(integerReplacement(((n-1)/2)+1)))
	}
	return 1 + min(integerReplacement(n+1), integerReplacement(n-1))
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	n := 2147483647
	fmt.Println(integerReplacement(n))
}
