package main

import (
	"fmt"
)

func reverseVowels(s string) string {
	stack := make([]byte, 0)
	queue := make([]int, 0)
	tmp := make([]byte, len(s))
	for i, c := range []byte(s) {
		tmp[i] = c
		if isVowels(c) {
			stack = append(stack, c)
			queue = append(queue, i)
		}
	}
	for ptr := 0; ptr < len(queue); ptr++ {
		vowel := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		tmp[queue[ptr]] = vowel
	}
	return string(tmp)
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
}
