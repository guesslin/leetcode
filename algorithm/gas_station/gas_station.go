package main

import (
	"fmt"
)

func canCompleteCircuit(gas []int, cost []int) int {
	diff := make([]int, len(gas))
	sumOfGas, sumOfCost := 0, 0
	for i := range diff {
		diff[i] = gas[i] - cost[i] // gas - cost => how much gas left while arrive next gas station
		sumOfGas += gas[i]
		sumOfCost += cost[i]
	}
	if sumOfCost > sumOfGas {
		return -1
	}
	stations := len(gas)
	for i := 0; i < stations; i++ {
		restGas := 0
		for j := 0; j < stations; j++ {
			restGas += diff[(i+j)%stations]
			if restGas < 0 {
				break
			}
		}
		if restGas >= 0 {
			return i
		}
	}
	return -1
}

func main() {
	gas := []int{1, 2, 3, 3}
	cost := []int{2, 1, 5, 1}
	fmt.Println(gas)
	fmt.Println(cost)
	fmt.Println(canCompleteCircuit(gas, cost))
}
