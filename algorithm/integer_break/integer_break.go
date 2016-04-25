package main

import (
	"fmt"
)

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	if n == 4 {
		return 4
	}
	x := n / 3
	if n%3 == 0 {
		return pow(3, x)
	}
	if n%3 == 1 {
		return pow(3, x-1) * 4
	}
	return pow(3, x) * 2
}

func pow(x, p int) int {
	r := x
	for i := 1; i < p; i++ {
		r *= x
	}
	return r
}

func main() {
	for i := 2; i <= 13; i++ {
		fmt.Println(integerBreak(i))
	}
}
