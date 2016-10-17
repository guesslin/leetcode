package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	count, cur := 0, 0
	for i := len(A) - 1; i >= 0; i-- {
		if A[i] == 1 {
			cur++
		} else {
			count += cur
			if count > 1000000000 {
				return -1
			}
		}
	}
	return count
}
