package main

import (
	"fmt"
)

func reverseString(s string) string {
	r := make([]rune, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, rune(s[i]))
	}
	return string(r)
}

func main() {
	fmt.Println(reverseString("hello"))
}
