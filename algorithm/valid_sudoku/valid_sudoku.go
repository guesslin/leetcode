package main

import (
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
	size := 9
	var pos int
	// check row
	for i := 0; i < size; i++ {
		check := make([]int, size)
		for j := 0; j < size; j++ {
			if '1' <= board[i][j] && board[i][j] <= '9' {
				pos = int(board[i][j] - '1')
				check[pos]++
				if check[pos] > 1 {
					return false
				}
			}
		}
	}
	// check column
	for i := 0; i < size; i++ {
		check := make([]int, size)
		for j := 0; j < size; j++ {
			if '1' <= board[j][i] && board[j][i] <= '9' {
				pos = int(board[j][i] - '1')
				check[pos]++
				if check[pos] > 1 {
					return false
				}
			}
		}
	}
	// check each box
	for i := 0; i < size; i++ {
		check := make([]int, size)
		r := i % 3
		c := i / 3
		for rp := r * 3; rp <= r*3+2; rp++ {
			for cp := c * 3; cp <= c*3+2; cp++ {
				if '1' <= board[cp][rp] && board[cp][rp] <= '9' {
					pos = int(board[cp][rp] - '1')
					check[pos]++
					if check[pos] > 1 {
						return false
					}
				}
			}
		}
	}
	return true
}

func main() {
	/*
		t := [][]byte{
			[]byte(".87654321"),
			[]byte("2........"),
			[]byte("3........"),
			[]byte("4........"),
			[]byte("5........"),
			[]byte("6........"),
			[]byte("7........"),
			[]byte("8........"),
			[]byte("9........"),
		}
	*/
	t2 := [][]byte{
		[]byte("....5..1."),
		[]byte(".4.3....."),
		[]byte(".....3..1"),
		[]byte("8......2."),
		[]byte("..2.7...."),
		[]byte(".15......"),
		[]byte(".....2..."),
		[]byte(".2.9....."),
		[]byte("..4......"),
	}
	// fmt.Println(isValidSudoku(t))
	fmt.Println(isValidSudoku(t2))
}
