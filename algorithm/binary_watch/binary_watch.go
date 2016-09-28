package main

import (
	"fmt"
)

func readBinaryWatch(num int) []string {
	res := make([]string, 0)
	mapHour := make(map[int][]int)
	// generate all hour
	for i := 0; i < 12; i++ {
		oneBits := countOneBits(i)
		mapHour[oneBits] = append(mapHour[oneBits], i)
	}
	mapMinute := make(map[int][]int)
	for i := 0; i < 60; i++ {
		oneBits := countOneBits(i)
		mapMinute[oneBits] = append(mapMinute[oneBits], i)
	}
	for hour := 0; hour <= num; hour++ {
		minu := num - hour
		// we combine all possible hour and minutes into res
		for i := 0; i < len(mapHour[hour]); i++ {
			for j := 0; j < len(mapMinute[minu]); j++ {
				res = append(res, fmt.Sprintf("%d:%02d", mapHour[hour][i], mapMinute[minu][j]))
			}
		}
	}
	return res
}

func countOneBits(num int) int {
	count := 0
	for num > 0 {
		count += num & 1
		num = num >> 1
	}
	return count
}

func main() {
	fmt.Println(readBinaryWatch(0))
}
