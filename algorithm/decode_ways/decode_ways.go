package main

import (
	"fmt"
)

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	res := make([]int, len(s)+1)
	res[0], res[1] = 1, 1
	for i := 2; i <= len(s); i++ {
		p := int(s[i-2] - '0') // 前一位數
		c := int(s[i-1] - '0') // 自己這一位數
		if c != 0 {
			res[i] += res[i-1]
		}
		if p*10+c > 9 && p*10+c < 27 {
			res[i] += res[i-2]
		}
	}
	return res[len(s)]
}

func main() {
	s := "10"

	fmt.Println(s, numDecodings(s))
}
