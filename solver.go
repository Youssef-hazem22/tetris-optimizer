package main

func canPlace(board [][]rune, t Tetromino, row, col int) bool {
	n := len(board)
	for _, p := range t.Cells {
		r, c := row+p.Row, col+p.Col

		if r < 0 || r >= n || c < 0 || c >= n {
			return false
		}
		if board[r][c] != '.' {
			return false
		}
	}
	return true
}

func place(board [][]rune, t Tetromino, row, col int) {
	for _, p := range t.Cells {
		board[row+p.Row][col+p.Col] = t.Letter
	}
}

func remove(board [][]rune, t Tetromino, row, col int) {
	for _, p := range t.Cells {
		board[row+p.Row][col+p.Col] = '.'
	}
}

func solve(board [][]rune, pieces []Tetromino, i int) bool {
	if i == len(pieces) {
		return true
	}
	n := len(board)
	t := pieces[i]
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if !canPlace(board, t, row, col) {
				continue
			}
			place(board, t, row, col)
			if solve(board, pieces, i+1) {
				return true
			}
			remove(board, t, row, col)
		}

	}
	return false
}
