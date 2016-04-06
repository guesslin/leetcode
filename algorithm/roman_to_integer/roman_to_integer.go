// https://leetcode.com/problems/roman-to-integer/
// Given a roman numeral, convert it to an integer.
// Input is guaranteed to be within the range from 1 to 3999.
package main

import (
	"fmt"
)

func romanToInt(s string) int {
	roman := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	var current, next int
	for i := 0; i < len(s); i++ {
		c := roman[string(s[i])]
		if i <= len(s)-2 {
			next = roman[string(s[i+1])]
			if next > c {
				c = next - c
				i++
			}
		}
		current += c
	}
	return current
}

func main() {
	fmt.Println("Expect:\t621\tGet:\t", romanToInt("DCXXI"))
	fmt.Println("Expect:\t421\tGet:\t", romanToInt("CDXXI"))
	fmt.Println("Expect:\t3021\tGet:\t", romanToInt("MMMXXI"))
	fmt.Println("Expect:\t3333\tGet:\t", romanToInt("MMMCCCXXXIII"))
	fmt.Println("Expect:\t704\tGet:\t", romanToInt("DCCIV"))
}
