package main

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	area1 := (C - A) * (D - B)
	area2 := (G - E) * (H - F)
	overlap := 0
	lowerTop := min(D, H)
	upperBot := max(B, F)
	lefterRight := min(C, G)
	righterLeft := max(A, E)
	if lowerTop > upperBot && lefterRight > righterLeft {
		overlap = (lowerTop - upperBot) * (lefterRight - righterLeft)
	}
	return area1 + area2 - overlap
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
