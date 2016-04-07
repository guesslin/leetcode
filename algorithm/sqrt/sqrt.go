package main

import (
	"fmt"
	"math"
)

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	xn := x / 2
	for xn*xn > x {
		xn = (xn + (x / xn)) / 2
	}
	return xn
}

func main() {
	fmt.Printf("mysqrt is %d, math.Sqrt is %f\n", mySqrt(0), math.Sqrt(0))
	fmt.Printf("mysqrt is %d, math.Sqrt is %f\n", mySqrt(1), math.Sqrt(1))
	fmt.Printf("mysqrt is %d, math.Sqrt is %f\n", mySqrt(2), math.Sqrt(2))
	fmt.Printf("mysqrt is %d, math.Sqrt is %f\n", mySqrt(3), math.Sqrt(3))
	fmt.Printf("mysqrt is %d, math.Sqrt is %f\n", mySqrt(4), math.Sqrt(4))
	fmt.Printf("mysqrt is %d, math.Sqrt is %f\n", mySqrt(9), math.Sqrt(9))
	fmt.Printf("mysqrt is %d, math.Sqrt is %f\n", mySqrt(122), math.Sqrt(122))
	fmt.Printf("mysqrt is %d, math.Sqrt is %f\n", mySqrt(981), math.Sqrt(981))
	fmt.Printf("mysqrt is %d, math.Sqrt is %f\n", mySqrt(4096), math.Sqrt(4096))
	fmt.Printf("mysqrt is %d, math.Sqrt is %f\n", mySqrt(2147395599), math.Sqrt(2147395599))
}
