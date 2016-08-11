package main

import (
	"fmt"
	"strconv"
)

// countAndSay is to generate n-th element of Look-and-Say sequence
// Look-and-Say sequence is defined at https://en.wikipedia.org/wiki/Look-and-say_sequence
// the first 10 element is 1, 11, 21, 1211, 111221, 312211, 13112221, 1113213211, 31131211131221, 13211311123113112211
// n must >= 1
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	pre := countAndSay(n - 1)
	res := make([]byte, 0, len(pre))
	count := 1
	char := pre[0]
	for i := 1; i <= len(pre); i++ {
		if i == len(pre) {
			res = append(res, strconv.Itoa(count)...)
			res = append(res, char)
			break
		}
		if char == pre[i] {
			count++
			continue
		}
		res = append(res, strconv.Itoa(count)...)
		res = append(res, char)
		count = 1
		char = pre[i]
	}
	return string(res)
}

func main() {
	var nums int
	fmt.Printf("Give me a number: ")
	fmt.Scanf("%d\n", &nums)
	fmt.Println(countAndSay(nums))
}
