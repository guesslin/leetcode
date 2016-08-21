package main

import (
	"fmt"
	"strings"
)

func shortestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	rev := reverse([]byte(s))
	if compare(s, string(rev)) {
		return s
	}
	lps := make([]int, n)
	i, j := 1, 0
	for i < n {
		if rev[i] == s[j] {
			j++
			lps[i] = j
			i++
		} else {
			if j == 0 {
				i++
			} else {
				j = lps[j-1]
			}
		}
	}
	suf := []byte(s)[lps[len(lps)-1]:]
	return string(append(rev, suf...))
}

func compare(s1, s2 string) bool {
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func reverse(s []byte) []byte {
	tmp := make([]byte, len(s))
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		tmp[i], tmp[j] = s[j], s[i]
	}
	return tmp
}

func main() {
	testCases := []struct {
		input, expect string
	}{
		{input: "", expect: ""},
		{input: "aba", expect: "aba"},
		{input: "abcd", expect: "dcbabcd"},
		{input: "aacecaaa", expect: "aaacecaaa"},
		{input: "aacecaaabc", expect: "cbaaacecaaabc"},
		{input: "bcdaffa", expect: "affadcbcdaffa"},
	}
	for _, c := range testCases {
		res := shortestPalindrome(c.input)
		if strings.Compare(res, c.expect) != 0 {
			fmt.Printf("Input\t%s\nExpect\t%s\nGet\t%s\n\n", c.input, c.expect, res)
		}
	}
}
