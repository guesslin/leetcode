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

func countBattleships2(board [][]byte) int {
	ships := 0
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[x]); y++ {
			if board[x][y] == 'X' {
				ships++
				// fill x direction 'X's to '.'
				for p := x + 1; p < len(board); p++ {
					if board[p][y] == '.' {
						break
					}
					board[p][y] = '.'

				}
				// fill y directino 'X's to '.'
				for p := y + 1; p < len(board[x]); p++ {
					if board[x][p] == '.' {
						break
					}
					board[x][p] = '.'
				}
			}
		}
	}
	return ships
}
