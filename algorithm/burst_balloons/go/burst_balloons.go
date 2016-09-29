package main

import (
	"fmt"
)

func burstBalloons(nums []int) int {
	l := len(nums) + 2
	balloons := make([]int, l)
	balloons[0], balloons[l-1] = 1, 1
	for i := 0; i < l-2; i++ {
		balloons[i+1] = nums[i]
	}
	counts := make([][]int, 0, l)
	for i := 0; i < l; i++ {
		counts = append(counts, make([]int, l))
	}
	for distance := 2; distance < l; distance++ {
		for left := 0; left < l-distance; left++ {
			right := left + distance
			// pick up a balloons between left and right
			for i := left + 1; i < right; i++ {
				counts[left][right] = max(counts[left][right], balloons[left]*balloons[i]*balloons[right]+counts[left][i]+counts[i][right])
			}

		}
	}
	return counts[0][l-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	balloons := []int{3, 1, 5, 8}
	fmt.Println(balloons)
	fmt.Println(burstBalloons(balloons))
}
