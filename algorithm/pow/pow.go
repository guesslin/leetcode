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
	result := myPow(x*x, n/2)
	if n/2*2 != n {
		result = result * x
	}
	if flag == -1 {
		return 1 / result
	}
	return result
}

func main() {
	fmt.Println(myPow(2, 1))
	fmt.Println(myPow(2, 0))
	fmt.Println(myPow(2, 14))
	fmt.Println(myPow(0.1, -2))
	fmt.Println(myPow(0.1, 2))
	fmt.Println(myPow(8.88023, 3))
	fmt.Println(myPow(2, -2))
}
