package main

import (
	"fmt"
)

func addDigits(num int) int {
	if num < 10 {
		return num
	}
	result := num % 9
	if result == 0 {
		return 9
	}
	return result
}

func main() {
	fmt.Println(addDigits(123456789))
}
