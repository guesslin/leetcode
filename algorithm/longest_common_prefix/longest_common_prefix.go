package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := make([]byte, 0)
	minLen := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if minLen > len(strs[i]) {
			minLen = len(strs[i])
		}
	}
	for i := 0; i < minLen; i++ {
		common := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if common != strs[j][i] {
				return string(res)
			}
		}
		res = append(res, common)
	}
	return string(res)
}

func main() {
	strs := []string{"a"}
	fmt.Println(longestCommonPrefix(strs))
}
