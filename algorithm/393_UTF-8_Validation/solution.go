package main

import "fmt"

func main() {
	data := []int{197, 130, 1}
	fmt.Println(validUtf8(data))
}

func validUtf8(data []int) bool {
	i := 0
	for i < len(data) {
		steps := lsbUtf8(data[i])
		if steps < 0 {
			return false
		}
		if steps == 1 {
			i++
			continue
		}

		for n := 1; n < steps; n++ {
			if len(data) <= i+n || !isLsb(data[i+n]) {
				return false
			}
		}

		i = i + steps
	}
	return true
}

func isLsb(d int) bool {
	return 0xc0&d == 0x80
}

func lsbUtf8(n int) int {
	switch {
	case n&0x80 == 0:
		return 1
	case n&0xe0 == 0xc0:
		return 2
	case n&0xf0 == 0xe0:
		return 3
	case n&0xf8 == 0xf0:
		return 4
	}
	return -1
}
