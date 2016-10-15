package main

func findMaximumXOR(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			tmp := nums[i] ^ nums[j]
			if tmp > res {
				res = tmp
			}
		}
	}
	return res
}
