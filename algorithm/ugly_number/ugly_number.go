package main

import (
	"fmt"
)

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}
	dividers := []int{2, 3, 5}
	for _, d := range dividers {
		for num%d == 0 {
			num = num / d
		}
	}
	return num == 1
}

func main() {
	for i := 1; i <= 20; i++ {
		fmt.Printf("%d is ugly? %v\n", i, isUgly(i))
	}
}
