package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(X int, A []int) int {
	// write your code in Go 1.4
	needs := make(map[int]bool)
	for i := 1; i <= X; i++ {
		needs[i] = true
	}
	for i, k := range A {
		if _, ok := needs[k]; ok {
			delete(needs, k)
		}
		if len(needs) == 0 {
			return i
		}
	}
	return -1
}
