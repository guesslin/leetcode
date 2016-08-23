package main

import (
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {
	need := make([]int, 26)
	has := make([]int, 26)
	for i := 0; i < len(ransomNote); i++ {
		need[int(ransomNote[i]-'a')]++
	}
	for i := 0; i < len(magazine); i++ {
		has[int(magazine[i]-'a')]++
	}
	for i := 0; i < 26; i++ {
		if need[i] > has[i] {
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
