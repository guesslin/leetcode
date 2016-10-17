package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	records := make([]int, len(A))
	for _, c := range A {
		if c > len(A) {
			return 0
		}
		records[c-1]++
		if records[c-1] > 1 {
			return 0
		}
	}
	for _, c := range records {
		if c != 1 {
			return 0
		}
	}
	return 1
}
