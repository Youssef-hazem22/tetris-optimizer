package main

import "strings"

func minBoardSize(pieces []Tetromino) int {
	maxcols := 0
	for _, piece := range pieces {
		for _, cell := range piece.Cells {
			if cell.Col > maxcols {
				maxcols = cell.Col
			}
		}
	}
	if maxcols < 1 {
		return 0
	}
	return maxcols
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
