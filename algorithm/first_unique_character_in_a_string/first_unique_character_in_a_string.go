package main

import (
	"fmt"
)

func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}
	words := make([]int, 26)
	slow := 0
	n := len(s)
	for i := 0; i < n; i++ {
		w := int(s[i] - 'a')
		words[w]++
		if slow < n && words[int(s[slow]-'a')] > 1 {
			for ; slow < n && words[int(s[slow]-'a')] > 1; slow++ {
			}
		}
	}
	if slow >= n {
		return -1
	}
	return slow
}

func main() {
	str := "aadadaad"
	fmt.Println(str)
	fmt.Println(firstUniqChar(str))
}
