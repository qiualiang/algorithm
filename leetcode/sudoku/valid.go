package sudoku

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		rows := make(map[byte]bool)
		cols := make(map[byte]bool)
		blocks := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			//valid every row
			if v := board[i][j]; v != '.' {
				if rows[v] {
					return false
				}
				rows[v] = true
			}

			//valid every col
			if v := board[j][i]; v != '.' {
				if cols[v] {
					return false
				}
				cols[v] = true
			}

			//valid blocks
			row := (i%3)*3 + j%3
			col := (i/3)*3 + j/3
			if v := board[row][col]; v != '.' {
				if blocks[v] {
					return false
				}
				blocks[v] = true
			}

		}
	}

	return true

}
