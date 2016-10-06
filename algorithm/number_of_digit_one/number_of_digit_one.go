package main

func countDigitOne(n int) int {
	res, d, f := 0, 0, 1
	prev := 0
	for n > 0 {
		r := n % 10
		res += f / 10 * r * d
		if r == 1 {
			res += 1 + prev
		} else if r > 0 {
			res += f
		}
		prev += r * f
		d++
		f *= 10
		n /= 10
	}
	return res
}
