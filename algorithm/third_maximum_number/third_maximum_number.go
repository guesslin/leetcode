package main

func thirdMax(nums []int) int {
	minNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if minNum > nums[i] {
			minNum = nums[i]
		}
	}
	maxNum := make([]int, 3)
	for i := 0; i < 3; i++ {
		maxNum[i] = minNum
	}
	for i := 0; i < len(nums); i++ {
		comp := nums[i]
		for j := 0; j < 3; j++ {
			if comp > maxNum[j] {
				maxNum[j], comp = comp, maxNum[j]
			} else if comp == maxNum[j] {
				break
			}
		}
	}
	if maxNum[2] == maxNum[1] {
		return maxNum[0]
	}
	return maxNum[2]
}
