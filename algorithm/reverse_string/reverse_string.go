package main

import (
	"fmt"
)

func reverseString(s string) string {
	var r string
	for _, c := range []rune(s) {
		r = string(c) + r
	}
	return r
}

func main() {
	fmt.Println(reverseString("hello"))
}
