package main

func maxArea(height []int) int {
	lo, hi, water := 0, len(height)-1, 0
	for lo < hi {
		water = max(water, (hi-lo)*min(height[lo], height[hi]))
		if height[lo] < height[hi] {
			lo++
		} else {
			hi--
		}
	}
	return water
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
