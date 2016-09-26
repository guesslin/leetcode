package main

import (
	"fmt"
)

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for _, b := range []byte(s) {
		switch b {
		case '[':
			fallthrough
		case '(':
			fallthrough
		case '{':
			stack = append(stack, b)
		case ']':
			fallthrough
		case ')':
			fallthrough
		case '}':
			if len(stack) == 0 {
				return false
			}
			p := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if b-p > 2 {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	s := "[[]]()"

	fmt.Println(s, isValid(s))
}
