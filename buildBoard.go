package main

import (
	"math"
	"strings"
)

func minBoardSize(pieceCount int) int {
	cells := 4 * pieceCount
	return int(math.Ceil(math.Sqrt(float64(cells))))
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
