package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	dict := make(map[rune]int)
	for _, c := range s {
		dict[c] += 1
	}
	for _, c := range t {
		dict[c] -= 1
	}
	for _, i := range dict {
		if i != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}
