package main

import (
	"fmt"
)

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	flag := 1
	if n < 0 {
		flag = -1
		n = -n
	}
	if n/2*2 != n {
		if flag == -1 {
			return 1 / (myPow(x*x, n/2) * x)
		}
		return myPow(x*x, n/2) * x
	}
	if flag == -1 {
		return 1 / myPow(x*x, n/2)
	}
	return myPow(x*x, n/2)
}

func main() {
	fmt.Println(myPow(2, 1))
	fmt.Println(myPow(2, 14))
	fmt.Println(myPow(2, -2))
}
