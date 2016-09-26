package main

import (
	"fmt"
)

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	isNegative := false
	if num < 0 {
		isNegative = true
		num = num * -1
	}
	res := make([]uint8, 0)
	for num > 0 {
		p := num % 16
		res = append(res, uint8(p))
		num = num / 16
	}
	if !isNegative {
		return string(convertI2B(reverse(res)))
	}
	// convert to 1's complement
	for i := range res {
		res[i] = ^res[i] & 15
	}
	for i := len(res); i < 8; i++ {
		res = append(res, uint8(15))
	}
	// add one to convert to 2's complement
	res[0]++
	for i := 1; res[i-1]&240 != 0; i++ {
		res[i-1] = res[i-1] & 15
		res[i]++
	}
	res = reverse(res)
	return string(convertI2B(res))
}

func convertI2B(nums []uint8) []byte {
	res := make([]byte, len(nums))
	for i, c := range nums {
		c = c & 15
		if c < 10 {
			res[i] = '0' + byte(c)
		} else {
			res[i] = 'a' + byte(c-10)
		}
	}
	return res
}

func reverse(s []uint8) []uint8 {
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func main() {
	testCases := []struct {
		num int
	}{
		{num: 26},
		{num: 0},
		{num: -1},
		{num: -99},
		{num: -0},
		{num: 39},
		{num: 777},
	}
	for i, c := range testCases {
		fmt.Println(i, c.num, toHex(c.num))
	}

}
