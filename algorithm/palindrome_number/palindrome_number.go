package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	orix := x
	a := 0
	for x > 0 {
		a = a*10 + x%10
		x = x / 10
	}
	return a == orix
}

func main() {
	fmt.Println("11000", isPalindrome(11000))
	fmt.Println("131000", isPalindrome(131000))
	fmt.Println("10", isPalindrome(10))
	fmt.Println("1001", isPalindrome(1001))
	fmt.Println("10101", isPalindrome(10101))
	fmt.Println("999", isPalindrome(999))
	fmt.Println("2345", isPalindrome(2345))
	fmt.Println("998899", isPalindrome(998899))
	fmt.Println("-101", isPalindrome(-101))
	fmt.Println("-997", isPalindrome(-997))
}
