package main

import (
	"fmt"
)

func reverseVowels(s string) string {
	iter := []byte(s)
	for i, j := 0, len(iter)-1; i < j; i, j = i+1, j-1 {
		for i < j && !isVowels(iter[i]) {
			i++
		}
		for j > i && !isVowels(iter[j]) {
			j--
		}
		iter[i], iter[j] = iter[j], iter[i]
	}
	return string(iter)
}

func isVowels(b byte) bool {
	// 遇到母音就回傳 true, 剩下的 回傳 false
	if b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U' {
		return true
	}
	if b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' {
		return true
	}
	return false
}

func main() {
	fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels(" a"))
	fmt.Println(reverseVowels("Aa"))
	fmt.Println(reverseVowels(",."))
}
