package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	record := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if c, ok := record[s[i]]; ok && c != t[i] {
			return false
		} else {
			record[s[i]] = t[i]
		}
	}
	record2 := make(map[byte]byte)
	for i := 0; i < len(t); i++ {
		if c, ok := record2[t[i]]; ok && c != s[i] {
			return false
		} else {
			record2[t[i]] = s[i]
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
