package main

import (
	"fmt"
)

func isHappy(n int) bool {
	var record []int
	sum := 0
	for n != 1 {
		for n != 0 {
			sum += square(n % 10)
			n /= 10
		}
		n = sum
		sum = 0
		if found(record, n) {
			return false
		}
		record = append(record, n)
	}
	return true
}

func found(space []int, num int) bool {
	for _, item := range space {
		if item == num {
			return true
		}
	}
	return false
}

func square(n int) int {
	return n * n
}

func main() {
	var num int
	fmt.Scanf("%d\n", &num)
	fmt.Printf("%d is Happy Number? %v\n", num, isHappy(num))
}
