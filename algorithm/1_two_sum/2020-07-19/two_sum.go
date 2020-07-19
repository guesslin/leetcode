package main

import "fmt"

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
*/
/*
// This two sum use 72ms, 3MB
// Idea: two-cursor to iterate the array and compare if sum of two cursor can fit the target
// Time-complexity: O(n^2)
func twoSum(nums []int, target int) []int {
	for i, ptr1 := range nums {
		for j, ptr2 := range nums {
			if i == j {
				continue
			}
			if ptr1+ptr2 == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
*/

// This use 4ms, 3.9MB
// Idea: Store the number into a reverse-map with (value => index-of-original-array)
// Iterate the given array and fill up the reverse-map together.
// Time-complexity: O(n)/* iteration */ + O(n) /* map-lookup */ => O(n)
func twoSum(nums []int, target int) []int {
	records := make(map[int]int)
	for i, num := range nums {
		search := target - num
		if idx, ok := records[search]; ok {
			return []int{idx, i}
		}
		records[num] = i
	}
	return nil
}

func main() {
	fmt.Println("1. Two-sum Start")
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{0, 4, 3, 0}, 0))
}
