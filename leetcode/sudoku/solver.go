package sudoku

func solveSudoku(board [][]byte) {
	if board == nil {
		return
	}
	solve(board)

}

func solve(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				var c byte
				for c = '1'; c <= '9'; c++ {
					if isValid(board, i, j, c) {
						board[i][j] = c
						if solve(board) {
							return true
						} else {
							board[i][j] = '.'
						}
					}
				}
				return false
			}

		}
	}
	return true
}

func isValid(board [][]byte, row int, col int, c byte) bool {
	for i := 0; i < 9; i++ {
		//check row
		if board[row][i] != '.' && board[row][i] == c {
			return false
		}
		if board[i][col] != '.' && board[i][col] == c {
			return false
		}
		if board[(row/3)*3+i/3][(col/3)*3+i%3] != '.' && board[(row/3)*3+i/3][(col/3)*3+i%3] == c {
			return false
		}
	}
	return true
}
