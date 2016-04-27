package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	carry := 0
	if digits[0] >= 0 {
		digits[len(digits)-1] += 1
		carry = digits[len(digits)-1] / 10
		digits[len(digits)-1] = digits[len(digits)-1] % 10
		for i := len(digits) - 2; i >= 0 && carry != 0; i-- {
			digits[i] = digits[i] + carry
			carry = digits[i] / 10
			digits[i] = digits[i] % 10
		}
		if carry > 0 {
			digits = append([]int{carry}, digits...)
		}
	} else {
		digits[len(digits)-1] = digits[len(digits)-1] - 1
		digits[0] = -digits[0]
		if digits[len(digits)-1] < 0 {
			carry = 1
			digits[len(digits)-1] = 9
		}
		for i := len(digits) - 2; i >= 0 && carry != 0; i-- {
			digits[i] = digits[i] - carry
			carry = 0
			if digits[i] < 0 {
				carry = 1
				digits[i] = 10 + digits[i]
			}
		}
		if digits[0] == 0 {
			digits = digits[1:]
		}
		digits[0] = -digits[0]
	}
	return digits
}

func main() {
	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{1}))
	fmt.Println(plusOne([]int{9, 9, 9}))
	fmt.Println(plusOne([]int{-1, 2, 0}))
	fmt.Println(plusOne([]int{-1, 0}))
	fmt.Println(plusOne([]int{-2, 0}))
}
