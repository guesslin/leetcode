package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}
	failure := make([]int, len(needle))
	failure[0] = -1
	for i, j := 1, failure[0]; i < len(needle); i++ {
		for j >= 0 && needle[j+1] != needle[i] {
			j = failure[j]
		}
		if needle[j+1] == needle[i] {
			j++
		}
		failure[i] = j
	}
	for i, j := 0, 0; i < len(haystack) && j < len(needle); {
		if j == len(needle)-1 && needle[j] == haystack[i] {
			fmt.Println(i, j)
			return i - j
		}
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			i = i - j + failure[j] + 2
			j = 0
		}
	}
	return -1
}

func main() {
	haystack := "astrstringstr"
	needle := "strstring"
	fmt.Println(strStr(haystack, needle))
}
