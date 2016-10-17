package main

func countBattleships(board [][]byte) int {
	ships := 0
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[x]); y++ {
			if board[x][y] == 'X' {
				if x-1 >= 0 && board[x-1][y] == 'X' {
					// already counted
					continue
				}
				if y-1 >= 0 && board[x][y-1] == 'X' {
					// already counted
					continue
				}
				ships++
			}
		}
	}
	return ships
}
