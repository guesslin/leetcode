package main

import (
	"fmt"
)

func canCompleteCircuit(gas []int, cost []int) int {
	lo, hi, sum := 0, len(gas)-1, 0
	for lo < hi {
		if sum > 0 {
			sum += gas[lo] - cost[lo]
			lo++
		} else {
			sum += gas[hi] - cost[hi]
			hi--
		}
	}
	sum += gas[lo] - cost[lo]
	if sum < 0 {
		return -1
	}
	if gas[lo]-cost[lo] >= 0 {
		return lo
	}
	return lo + 1
}

func main() {
	gas := []int{1, 2, 3, 3}
	cost := []int{2, 1, 5, 1}
	fmt.Println(gas)
	fmt.Println(cost)
	fmt.Println(canCompleteCircuit(gas, cost))
}
