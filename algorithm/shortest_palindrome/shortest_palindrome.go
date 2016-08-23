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
	ori := []byte(s)
	if compare(s, string(rev)) {
		return s
	}
	// build longest palindromic substring using Manacher's Algorithm
	N := 2*n + 1
	lps := make([]int, N)
	lps[1] = 1
	c, r := 0, 0
	for i := 2; i < N; i++ {
		mirror := 2*c - i
		if i < r {
			lps[i] = min(r-i, lps[mirror])
		}
		for lb, rb := i-lps[i]-1, i+lps[i]+1; (lb >= 0 && rb < N && rb/2 < n) && ((rb%2 == 0) || ori[(rb)/2] == ori[(lb)/2]); lb, rb = i-lps[i]-1, i+lps[i]+1 {
			lps[i]++
		}
		if i+lps[i] > r {
			c = i
			r = lps[i] + i
		}
	}
	// take longest palindromic substring from start
	suffix := 0
	for i := 1; i < N; i++ {
		if lps[i] == i && i > suffix {
			suffix = i
		}
	}
	suf := []byte(s)[suffix:]
	return string(append(rev, suf...))
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
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
