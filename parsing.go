package main

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
