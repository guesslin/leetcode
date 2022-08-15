package main

import (
	"fmt"
	"math"
	"time"
)

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 49, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

func sqrt(in int) int {
	return int(math.Ceil(math.Sqrt(float64(in))))
}

func isPrime(t int) bool {
	if t <= 1 {
		return false
	}
	if t%2 == 0 {
		return false
	}
	for i := 3; i < sqrt(t); i += 2 {
		if t%i == 0 {
			return false
		}
	}
	return true
}

func checkPerfectNumber(num int) bool {
	if num <= 0 {
		return false
	}
	sum := 0
	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			sum += i
			if i*i != num {
				sum += num / i
			}
		}

	}
	return sum-num == num
}

func trigger(target int) {
	s := time.Now()
	fmt.Printf("Is %d a perfect number? %+v\n", target, checkPerfectNumber(target))
	e := time.Now()
	fmt.Printf("Execution time is %+v\n", e.Sub(s))
}

func main() {
	trigger(28)
	trigger(99999994)
}
