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
	x := n / 3
	remainder := n % 3
	if remainder == 0 {
		return pow(3, x)
	} else if remainder == 1 {
		return pow(3, x-1) * 4
	} else {
		return pow(3, x) * 2
	}
}

func pow(x, n int) int {
	if n == 0 {
		return 1
	}
	result := pow(x*x, n/2)
	if n/2*2 != n {
		result = result * x
	}
	return result
}

func main() {
	for i := 2; i <= 13; i++ {
		fmt.Printf("Number %d can break into max product %d\n", i, integerBreak(i))
	}
}
