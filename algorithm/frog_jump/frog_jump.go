package main

import (
	"fmt"
)

func canCross(stones []int) bool {
	// start stone is always sit on zero, and it will always be jumpable
	if len(stones) == 1 {
		return true
	}
	stonesMap := make(map[int]map[int]bool)
	for _, p := range stones {
		stonesMap[p] = make(map[int]bool)
	}
	// if the first jump can't jump to the second stone, then it will always be false
	if _, ok := stonesMap[1]; !ok {
		return false
	}
	maxStone := 1
	stonesMap[1][1] = true
	for i := 1; i < len(stones); i++ {
		if stones[i] > maxStone {
			break
		}
		for k, _ := range stonesMap[stones[i]] {
			if _, ok := stonesMap[stones[i]+k-1]; ok && (k-1) > 0 {
				stonesMap[stones[i]+k-1][k-1] = true
			}
			if _, ok := stonesMap[stones[i]+k]; ok {
				stonesMap[stones[i]+k][k] = true
			}
			if _, ok := stonesMap[stones[i]+k+1]; ok {
				stonesMap[stones[i]+k+1][k+1] = true
			}
			if stones[i]+k+1 > maxStone {
				maxStone = stones[i] + k + 1
			}
		}
	}
	return len(stonesMap[stones[len(stones)-1]]) > 0
}

func main() {
	stones := []int{0, 1, 3, 4, 5, 7, 9, 10, 12}
	fmt.Println(stones)
	if canCross(stones) {
		fmt.Println("Frog can cross the river!")
	} else {
		fmt.Println("Frog can NOT cross the river!")
	}
	stones2 := []int{0, 1, 3, 6, 10, 13, 15, 18}
	fmt.Println(stones2)
	if canCross(stones2) {
		fmt.Println("Frog can cross the river!")
	} else {
		fmt.Println("Frog can NOT cross the river!")
	}
}
