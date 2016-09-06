package main

import (
	"fmt"
)

func isPerfectSquare(num int) bool {
	if num <= 0 {
		return false
	}
	if num == 1 {
		return true
	}
	lo, hi := 1, num/2+1
	for lo <= hi {
		mid := (lo + hi) / 2
		if mid*mid == num {
			return true
		} else if mid*mid > num {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return false
}

func main() {
	for i := 0; i < 101; i++ {
		if isPerfectSquare(i) {
			fmt.Println(i, "isPerfectSquare")
		}
	}
}
