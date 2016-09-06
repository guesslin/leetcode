package main

import (
	"fmt"
)

func convertToTitle(n int) string {
	res := make([]rune, 0)
	for n > 0 {
		c := rune('A' + (n-1)%26)
		n = (n - 1) / 26
		res = append(res, c)
	}
	for i, j := 0, len(res)-1; i <= j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func main() {
	for i := 1; i < 100; i++ {
		fmt.Println(i, "=>", convertToTitle(i))
	}
}
