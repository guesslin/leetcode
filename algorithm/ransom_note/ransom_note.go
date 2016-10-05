package main

import (
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {
	have := make([]int, 26)
	for _, c := range []byte(magazine) {
		have[int(c-'a')]++
	}
	for _, c := range []byte(ransomNote) {
		have[int(c-'a')]--
		if have[int(c-'a')] < 0 {
			return false
		}
	}
	return true
}

func main() {
	testCases := []struct {
		ransomNote, magazine string
		result               bool
	}{
		{ransomNote: "", magazine: "", result: true},
		{ransomNote: "a", magazine: "b", result: false},
		{ransomNote: "aa", magazine: "ab", result: false},
		{ransomNote: "aa", magazine: "aab", result: true},
	}
	for _, c := range testCases {
		res := canConstruct(c.ransomNote, c.magazine)
		if res != c.result {
			fmt.Printf("ransom: %s, magazine: %s, expect: %v, get: %v\n", c.ransomNote, c.magazine, c.result, res)
		}
	}
}
