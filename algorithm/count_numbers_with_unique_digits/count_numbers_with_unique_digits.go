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
	last, cache, prod := 0, 0, 1
	for i := 1; i <= n; i++ {
		cache = i * prod
		cache += last
		last = cache
		prod *= (10 - i)
	}
	return 10*prod + cache
}

func main() {
	for i := 0; i < 11; i++ {
		fmt.Println(countNumbersWithUniqueDigits(i))
	}
}
