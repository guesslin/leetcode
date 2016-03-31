package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	orix := x
	a := 0
	for x > 0 {
		a = a*10 + x%10
		x = x / 10
		if a == x {
			return true
		}
	}
	if a == orix {
		return true
	}
	return false
}

func main() {
	fmt.Println("1001", isPalindrome(1001))
	fmt.Println("10101", isPalindrome(10101))
	fmt.Println("999", isPalindrome(999))
	fmt.Println("2345", isPalindrome(2345))
	fmt.Println("998899", isPalindrome(998899))
	fmt.Println("-101", isPalindrome(-101))
	fmt.Println("-997", isPalindrome(-997))
}
