package main

import (
	"fmt"
	"strconv"
	"strings"
)

func myAtoi(str string) int {
	res := 0
	signed := 1
	newStr := strings.TrimSpace(str)
	for i, c := range newStr {
		if i == 0 {
			if c == '-' {
				signed = -1
				continue
			} else if c == '+' {
				continue
			}
		}
		if '0' <= c && c <= '9' {
			res = res*10 + int(c-'0')
		} else {
			return signed * res
		}
		if res > 2147483647 {
			break
		}
	}
	if res > 2147483647 {
		if signed == 1 {
			res = 2147483647
		} else {
			res = 2147483648
		}
	}
	return signed * res
}

func main() {
	testStrings := []string{
		"+1",
		"123",
		"    010", // expect 10, but lib got 0
		"-100",
		"1e3",
		"jjj",
		"中文",
		"1.1",
		"9999.9",
		"a+1dfs中文",
		"  -0012a42", // expect -12, but lib got 0
		"2147483648", // leetcode expect 2147483647
		"9223372036854775809",
	}
	for _, s := range testStrings {
		standard, _ := strconv.Atoi(s)
		fmt.Println("string:", s, "\nLibary:", standard, "\nmyAtoi:", myAtoi(s))
		fmt.Println("=====================Case================================")
	}
}
