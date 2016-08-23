package main

import (
	"fmt"
)

type Uniq struct {
	Char  byte
	Index int
}

func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}
	words := make([]int, 26)
	first := make([]Uniq, 0, 26)
	n := len(s)
	for i := 0; i < n; i++ {
		w := int(s[i] - 'a')
		words[w]++
		if words[w] == 1 {
			first = append(first, Uniq{Char: s[i], Index: i})
		} else {
			for f, unit := range first {
				if unit.Char == s[i] {
					first = append(first[:f], first[f+1:]...)
					break
				}
			}
		}
	}
	if len(first) != 0 {
		return first[0].Index
	}
	return -1
}

func main() {
	str := "aadadaad"
	fmt.Println(str)
	fmt.Println(firstUniqChar(str))
}
