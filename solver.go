package main

import "strings"

func canPlace(board [][]byte, t Tetromino, row, col int) bool {
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

func place(board [][]byte, t Tetromino, row, col int) {
	for _, p := range t.Cells {
		board[row+p.Row][col+p.Col] = byte(t.Letter)
	}
}

func remove(board [][]byte, t Tetromino, row, col int) {
	for _, p := range t.Cells {
		board[row+p.Row][col+p.Col] = '.'
	}
}

func sameShape(a, b Tetromino) bool {
	for i := range a.Cells {
		if a.Cells[i] != b.Cells[i] {
			return false
		}
	}
	return true
}

func solve(board [][]byte, pieces []Tetromino, used []bool, idx, left, waste int) bool {
	if left == 0 {
		return true
	}
	n := len(board)

	for idx < n*n && board[idx/n][idx%n] != '.' {
		idx++
	}
	if idx >= n*n {
		return false
	}
	row, col := idx/n, idx%n

	for i, t := range pieces {
		if used[i] {
			continue
		}

		duplicate := false
		for j := 0; j < i; j++ {
			if !used[j] && sameShape(pieces[j], t) {
				duplicate = true
				break
			}
		}
		if duplicate {
			continue
		}
		pr, pc := row-t.Cells[0].Row, col-t.Cells[0].Col
		if pc < 0 || !canPlace(board, t, pr, pc) {
			continue
		}

		place(board, t, pr, pc)
		used[i] = true
		if solve(board, pieces, used, idx+1, left-1, waste) {
			return true
		}
		used[i] = false
		remove(board, t, pr, pc)
	}

	// option 2: leave this cell empty
	if waste > 0 {
		return solve(board, pieces, used, idx+1, left, waste-1)
	}
	return false
}

func minBoardSize(pieces []Tetromino) int {
	n := 2
	for n*n < 4*len(pieces) {
		n++
	}
	for _, piece := range pieces {
		for _, cell := range piece.Cells {
			if cell.Col+1 > n {
				n = cell.Col + 1
			}
			if cell.Row+1 > n {
				n = cell.Row + 1
			}
		}
	}
	return n
}

func newBoard(n int) [][]byte {
	board := make([][]byte, n)
	for r := range board {
		board[r] = make([]byte, n)
		for c := range board[r] {
			board[r][c] = '.'
		}
	}
	return board
}

func boardString(board [][]byte) string {
	var sb strings.Builder
	for _, row := range board {
		for _, cell := range row {
			sb.WriteByte(cell)
		}
		sb.WriteByte('\n')

	}
	return sb.String()
}
