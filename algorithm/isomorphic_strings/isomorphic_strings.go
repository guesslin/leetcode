package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	mapping := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if mapping[s[i]] == 0 {
			c := t[i]
			for _, n := range mapping {
				// if c is already mapping to other char
				if c == n {
					return false
				}
			}
			mapping[s[i]] = c
		} else if mapping[s[i]] != t[i] {
			return false
		}
	}
	return true
}

func main() {
	testCases := []struct {
		s, t string
		r    bool
	}{
		{s: "aa", t: "ab", r: false},
		{s: "ab", t: "aa", r: false},
	}
	for _, c := range testCases {
		res := isIsomorphic(c.s, c.t)
		if res != c.r {
			fmt.Printf("s: %s, t: %s, expect: %v, Get: %v\n", c.s, c.t, c.r, res)
		}
	}
}
