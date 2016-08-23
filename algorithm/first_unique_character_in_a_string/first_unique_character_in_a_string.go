package main

import (
	"fmt"
)

func firstUniqChar(s string) int {
	words := make([]int, 26)
	for _, c := range s {
		words[int(c)-'a']++
	}
	for i, c := range s {
		if words[int(c)-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	str := "leetcode"
	fmt.Println(str)
	fmt.Println(firstUniqChar(str))
}
