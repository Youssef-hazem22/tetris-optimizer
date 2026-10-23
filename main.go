package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strings"
)

const maxPieces = 26

type Point struct {
	Row, Col int
}

type Tetromino struct {
	Cells  []Point
	Letter rune
}

func readFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func splitBlocks(content string) ([][]string, error) {
	var blocks [][]string
	var current []string

	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			blocks = append(blocks, current)
			current = nil
			continue
		}
		current = append(current, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(current) > 0 {
		blocks = append(blocks, current)
	}
	return blocks, nil
}
func validateBlocks(blocks [][]string) error {
	if len(blocks) == 0 {
		return errors.New("no tetrominoes found")
	}
	if len(blocks) > maxPieces {
		return fmt.Errorf("too many tetrominoes: %d (max %d)", len(blocks), maxPieces)
	}

	for i, block := range blocks {
		if err := validateBlock(block); err != nil {
			return fmt.Errorf("block %d: %w", i+1, err)
		}
	}
	return nil
}

func validateBlock(block []string) error {
	if len(block) != 4 {
		return fmt.Errorf("expected 4 lines, got %d", len(block))
	}

	for r, line := range block {
		if len(line) != 4 {
			return fmt.Errorf("line %d: expected 4 characters, got %d", r+1, len(line))
		}
		for c, ch := range line {
			if ch != '#' && ch != '.' {
				return fmt.Errorf("line %d, col %d: invalid character %q", r+1, c+1, ch)
			}
		}
	}
	return nil
}

func parseTetrominoes(blocks [][]string) []Tetromino {
	pieces := make([]Tetromino, 0, len(blocks))
	for _, block := range blocks {
		pieces = append(pieces, parseTetromino(block))
	}
	return pieces
}

func parseTetromino(block []string) Tetromino {
	var cells []Point
	for r, line := range block {
		for c, ch := range line {
			if ch == '#' {
				cells = append(cells, Point{Row: r, Col: c})
			}
		}
	}
	return Tetromino{Cells: cells}
}
func validateShapes(pieces []Tetromino) error {
	for i, p := range pieces {
		if err := validateShape(p); err != nil {
			return fmt.Errorf("piece %d: %w", i+1, err)
		}
	}
	return nil
}

func validateShape(t Tetromino) error {
	if len(t.Cells) != 4 {
		return fmt.Errorf("expected 4 '#', got %d", len(t.Cells))
	}

	set := make(map[Point]bool, len(t.Cells))
	for _, p := range t.Cells {
		set[p] = true
	}

	directions := []Point{{Row: -1, Col: 0}, {Row: 1, Col: 0}, {Row: 0, Col: -1}, {Row: 0, Col: 1}}
	adjacencies := 0
	for _, p := range t.Cells {
		for _, d := range directions {
			neighbor := Point{Row: p.Row + d.Row, Col: p.Col + d.Col}
			if set[neighbor] {
				adjacencies++
			}
		}
	}

	if adjacencies != 6 && adjacencies != 8 {
		return fmt.Errorf("cells are not connected (adjacencies = %d)", adjacencies)
	}
	return nil
}
func normalizeAll(pieces []Tetromino) []Tetromino {
	result := make([]Tetromino, 0, len(pieces))
	for _, p := range pieces {
		result = append(result, normalize(p))
	}
	return result
}

func normalize(t Tetromino) Tetromino {
	minRow, minCol := t.Cells[0].Row, t.Cells[0].Col
	for _, p := range t.Cells {
		if p.Row < minRow {
			minRow = p.Row
		}
		if p.Col < minCol {
			minCol = p.Col
		}
	}

	cells := make([]Point, len(t.Cells))
	for i, p := range t.Cells {
		cells[i] = Point{Row: p.Row - minRow, Col: p.Col - minCol}
	}
	return Tetromino{Cells: cells, Letter: t.Letter}
}
func assignLetters(pieces []Tetromino) []Tetromino {
	result := make([]Tetromino, len(pieces))
	for i, p := range pieces {
		result[i] = Tetromino{
			Cells:  p.Cells,
			Letter: 'A' + rune(i),
		}
	}
	return result
}

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
	n := len(pieces)
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

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR")
		return
	}

	content, err := readFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	blocks, err := splitBlocks(content)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	if err := validateBlocks(blocks); err != nil {
		fmt.Println("ERROR")
		return
	}

	pieces := parseTetrominoes(blocks)

	if err := validateShapes(pieces); err != nil {
		fmt.Println("ERROR")
		return
	}

	pieces = normalizeAll(pieces)
	pieces = assignLetters(pieces)
	n := minBoardSize(len(pieces))
	for {
		board := newBoard(n)
		if solve(board, pieces, 0) {
			fmt.Print(boardString(board))
			return
		}
		n++
	}

}
