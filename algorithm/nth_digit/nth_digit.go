package main

import (
	"fmt"
)

// findNthDigit => 9 * 10^(L-1) * L
func findNthDigit(n int) int {
	if n < 10 {
		return n
	}
	l := 1
	base := 9
	for ; n > base*l; l, base = l+1, base*10 {
		n -= base * l
	}
	base = base / 9
	a := n / l
	b := n % l
	if b > 0 {
		a++
		b = l - b
	}
	res := base + a - 1
	fmt.Println(res)
	for i := 0; i < b; i++ {
		res = res / 10
	}
	return res % 10
}

func main() {
	/*
		for i := 1; i <= 20; i++ {
			fmt.Println(i, findNthDigit(i))
		}
	*/
	i := 1000
	fmt.Println(i, findNthDigit(i))
}
