package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	lower := toLower(s)
	for i, j := 0, len(lower)-1; i <= j; i, j = i+1, j-1 {
		if lower[i] != lower[j] {
			return false
		}
	}
	return true
}

func toLower(s string) []byte {
	res := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if !isAlphaNumeric(s[i]) {
			continue
		}
		if 'A' <= s[i] && s[i] <= 'Z' {
			res = append(res, s[i]-'A'+'a')
		} else {
			res = append(res, s[i])
		}
	}
	return res
}

func isAlphaNumeric(b byte) bool {
	if 'a' <= b && b <= 'z' {
		return true
	}
	if 'A' <= b && b <= 'Z' {
		return true
	}
	if '0' <= b && b <= '9' {
		return true
	}
	return false
}

func main() {
	testCases := []struct {
		input  string
		expect bool
	}{
		{input: "", expect: true},
		{input: "aa", expect: true},
		{input: "12345ABC", expect: false},
		{input: "12345:ABC:cba:54321", expect: true},
		{input: "AbcBa", expect: true},
		{input: "aba", expect: true},
		{input: " a", expect: true},
		{input: " abc a", expect: false},
		{input: "  abb: as, a bb-a", expect: true},
	}
	for _, c := range testCases {
		res := isPalindrome(c.input)
		if res != c.expect {
			fmt.Printf("Wrong while running \"%s\", expect %v, get %v\n", c.input, c.expect, res)
		} else {
			fmt.Printf("Correct \"%s\", %v\n", c.input, res)
		}
	}
}
