package main

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	store := make([]int, len(A))
	store[0] = A[0]
	for i := 1; i < len(A); i++ {
		store[i] = store[i-1]
		for k := 2; k <= 6; k++ {
			if i >= k {
				store[i] = max(store[i], store[i-k])
			} else {
				break
			}
		}
		store[i] += A[i]
	}
	return store[len(store)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
