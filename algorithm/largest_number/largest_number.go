package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type byNumSeq []int

func (p byNumSeq) Len() int {
	return len(p)
}

func (p byNumSeq) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p byNumSeq) Less(i, j int) bool {
	a := strconv.Itoa(p[i])
	b := strconv.Itoa(p[j])
	ab, _ := strconv.Atoi(strings.Join([]string{a, b}, ""))
	ba, _ := strconv.Atoi(strings.Join([]string{b, a}, ""))
	return ab < ba
}

func largestNumber(nums []int) string {
	sort.Sort(sort.Reverse(byNumSeq(nums)))
	fmt.Println(nums)
	text := make([]string, 0, len(nums))
	for i := range nums {
		text = append(text, strconv.Itoa(nums[i]))
	}
	res := strings.Join(text, "")
	// special case for all input are 0 like, [0, 0]
	// get "00" as result, have convert it back to normal zero
	if res[0] == '0' {
		res = "0"
	}
	return res
}

func main() {
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
	fmt.Println(largestNumber([]int{0, 0}))
}
