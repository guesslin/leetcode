package main

import (
	"fmt"
)

func rob(nums []int) int {
	houses := len(nums)
	// there is no house to rob
	if houses == 0 {
		return 0
	}
	// there is only one house to rob
	if houses == 1 {
		return nums[0]
	}
	headSelected := make([]bool, houses)
	lastSelected := false
	count := make([]int, houses)
	count[0] = nums[0]
	headSelected[0] = true
	if nums[0] > nums[1] {
		count[1] = nums[0]
		headSelected[1] = headSelected[0]
	} else {
		count[1] = nums[1]
	}
	for i := 2; i < houses; i++ {
		if count[i-1] > nums[i]+count[i-2] {
			count[i] = count[i-1]
			headSelected[i] = headSelected[i-1]
			lastSelected = false
		} else {
			count[i] = nums[i] + count[i-2]
			headSelected[i] = headSelected[i-2]
			lastSelected = true
		}
	}
	// 頭尾沒有相鄰
	if !headSelected[houses-1] || !lastSelected {
		return count[houses-1]
	}
	// 頭尾有相鄰，找減去頭或尾之後哪個數字大
	// 還要注意跟最後一個數字相比較
	if nums[0] > nums[houses-1] {
		count[houses-1] = count[houses-1] - nums[houses-1]
	} else {
		count[houses-1] = count[houses-1] - nums[0]
	}
	lastMax := 0
	if count[houses-1] > count[houses-2] {
		lastMax = count[houses-1]
	} else {
		lastMax = count[houses-2]
	}
	count[1] = nums[1]
	if nums[1] > nums[2] {
		count[2] = nums[1]
	} else {
		count[2] = nums[2]
	}
	for i := 3; i < houses; i++ {
		if count[i-1] > nums[i]+count[i-2] {
			count[i] = count[i-1]
		} else {
			count[i] = nums[i] + count[i-2]
		}
	}
	if lastMax > count[houses-1] {
		return lastMax
	}
	return count[houses-1]
}

func main() {
	//	fmt.Println(rob([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(rob([]int{2, 2, 4, 3, 2, 5}))
	//	fmt.Println(rob([]int{2, 1}))
	//	fmt.Println(rob([]int{2, 3, 2}))
	//	fmt.Println(rob([]int{}))
	//	fmt.Println(rob([]int{1, 2, 3, 4, 5, 6, 7}))
	//	fmt.Println(rob([]int{0}))
	//	fmt.Println(rob([]int{2, 1}))
	//	fmt.Println(rob([]int{2, 6, 1}))
}
