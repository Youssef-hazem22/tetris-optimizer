package main

import (
	"errors"
	"fmt"
)

func Allmatching(pieces []Tetromino) ([]Tetromino, error) {
	for _, p := range pieces {
		if err := onematching(p); err != nil {
			return nil, err
		}
	}
	return normalizeAll(pieces), nil
}

func onematching(t Tetromino) error {
	if len(t.Cells) != 4 {
		return errors.New("ERROR length")
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
		result = append(result, onenormalize(p))
	}
	return result
}

func onenormalize(t Tetromino) Tetromino {
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
