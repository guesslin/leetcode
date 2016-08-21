package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	mapping := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if mapping[s[i]] == 0 {
			mapping[s[i]] = t[i]
		} else if mapping[s[i]] != t[i] {
			return false
		}
	}
	mapping2 := make(map[byte]byte)
	for i := 0; i < len(t); i++ {
		if mapping2[t[i]] == 0 {
			mapping2[t[i]] = s[i]
		} else if mapping2[t[i]] != s[i] {
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
