package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	if n&(n-1) == 0 {
		return true
	}
	return false
}

func main() {
	for i := 0; i <= 536870912; i++ {
		if isPowerOfTwo(i) {
			fmt.Printf("%d is power of two\n", i)
		}
	}
}
