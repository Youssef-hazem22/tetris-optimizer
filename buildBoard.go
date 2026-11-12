package main

import (
	"math"
	"strings"
)

func minBoardSize(pieceCount int) int {
	cells := 4 * pieceCount
	return int(math.Ceil(math.Sqrt(float64(cells))))
}

func newBoard(n int) [][]rune {
	board := make([][]rune, n)
	for r := range board {
		board[r] = make([]rune, n)
		for c := range board[r] {
			board[r][c] = '.'
		}
	}
	return board
}

func boardString(board [][]rune) string {
	var sb strings.Builder
	for _, row := range board {
		for _, cell := range row {
			sb.WriteRune(cell)
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}
