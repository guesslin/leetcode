package main

import (
	"fmt"
)

func isHappy(n int) bool {
	record := make(map[int]bool)
	for n != 1 {
		sum := 0
		for n != 0 {
			sum += square(n % 10)
			n /= 10
		}
		n = sum
		if _, ok := record[n]; ok {
			return false
		}
		record[n] = true
	}
	return true
}

func square(n int) int {
	return n * n
}

func main() {
	var num int
	fmt.Scanf("%d\n", &num)
	fmt.Printf("%d is Happy Number? %v\n", num, isHappy(num))
}
