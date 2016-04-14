package main

import (
	"fmt"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	d := 2*numRows - 2
	buf := make([]string, numRows)
	var result string
	for i, c := range s {
		b := i % d
		if b < numRows {
			buf[b] += string(c)
		} else {
			b = b - numRows
			buf[numRows-2-b] += string(c)

		}
	}
	fmt.Println(buf)
	for _, c := range buf {
		result += c
	}
	return result
}

func main() {
	s := "PAYPALISHIRING"
	fmt.Println(s, convert(s, 3))
	fmt.Println(s, "PAHNAPLSIIGYIR")
	fmt.Println("A", convert("A", 1))
}
