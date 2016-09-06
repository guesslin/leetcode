package main

import (
	"fmt"
)

func isPerfectSquare(num int) bool {
	level := 1
	for num > 0 {
		num -= level
		level += 2
	}
	return num == 0
}

func main() {
	for i := 0; i < 101; i++ {
		fmt.Println(i, "isPerfectSquare:", isPerfectSquare(i))
	}
}
