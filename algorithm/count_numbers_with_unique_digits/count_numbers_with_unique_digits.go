package main

import (
	"fmt"
)

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n > 9 {
		n = 9
	}
	res := make([]int, n+1)
	res[0] = 10
	p := 9
	sum := 10
	for i := 1; i <= n; i++ {
		res[i] = p * (10 - i)
		p *= (10 - i)
		sum += res[i]
	}
	return sum
}

func main() {
	for i := 0; i < 11; i++ {
		fmt.Println(countNumbersWithUniqueDigits(i))
	}
}
