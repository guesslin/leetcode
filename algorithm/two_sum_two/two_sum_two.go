package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		j := bisearch(numbers[i+1:], target-numbers[i])
		if j != -1 {
			return []int{i + 1, i + j + 2}
		}
	}
	return nil
}

func bisearch(data []int, target int) int {
	lo, hi := 0, len(data)-1
	for lo <= hi {
		half := (lo + hi) / 2
		if data[half] == target {
			return half
		} else if data[half] > target {
			hi = half - 1
		} else {
			lo = half + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
