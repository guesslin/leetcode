package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	records := make(map[int]int)
	for _, c := range A {
		records[c]++
	}
	i := 1
	for ; len(records) > 0; i++ {
		if _, ok := records[i]; ok {
			delete(records, i)
		} else {
			return i
		}
	}
	return i
}
