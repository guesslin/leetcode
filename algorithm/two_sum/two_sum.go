package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	record := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := record[nums[i]]; ok {
			record[nums[i]] = append(record[nums[i]], i)
		} else {
			record[nums[i]] = []int{i}
		}
	}
	for i := 0; i < len(nums); i++ {
		if j, ok := record[target-nums[i]]; ok {
			for _, c := range j {
				if c != i {
					return []int{i, c}
				}
			}
		}
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{0, 4, 3, 0}, 0))
}
