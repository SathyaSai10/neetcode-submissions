func isValidSudoku(board [][]byte) bool {
    var rows, cols, boxes [9]map[byte]bool

	for i := range rows {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for i := 0; i < 81; i++ { // single flat loop, no nesting
		row, col := i/9, i%9
		val := board[row][col]

		if val == '.' {
			continue
		}

		box := (row/3)*3 + col/3

		if rows[row][val] || cols[col][val] || boxes[box][val] {
			return false
		}

		rows[row][val] = true
		cols[col][val] = true
		boxes[box][val] = true
	}

	return true
}
